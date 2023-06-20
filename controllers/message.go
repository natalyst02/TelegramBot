package controllers

import (
	"encoding/json"
	"fmt"
	"gitbot/configs"
	"gitbot/models"
	"gitbot/models/webhook"
	"gitbot/models/webhook/comment"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type msgResult struct {
	content, text, url string
	err                error
}

// var count int
var currentMesID int
var sendActions = map[string]func(body []byte) msgResult{
	"push":          pushMessage,
	"issue":         issueMessage,
	"note":          commentMessage,
	"merge_request": mergeRequestMessage,
	"wiki_page":     wikiPageMessage,
	"tag_push":      tagPushMessage,
	"feature_flag":  featFlagMessage,
	"deployment":    deployMessage,
	"release":       releaseMessage,
	// "pipeline":      pipelineMessage,
	// "build":         buildMessage,
}
var editActions = map[string]func(message tgbotapi.Message, body []byte) string{
	"build": buildAction,
	// "pipeline": pipelineAction,
}

var buildEmoji = map[string]string{
	"test":   "🧪",
	"build":  "⚙️",
	"deploy": "🚀",
}
var statusEmoji = map[string]string{
	"pending": "⌛",
	"running": "⏩",
	"success": "✅",
	"created": "✨",
	"failed":  "🚫",
}

func pushMessage(body []byte) msgResult {
	var p webhook.PushEventPayload
	err := json.Unmarshal(body, &p)
	if err != nil {
		log.Panicln(err)
	}
	BranchName := p.Ref
	s := strings.Split(BranchName, "heads/")
	var res string
	if len(p.Commits) == 0 {
		res = fmt.Sprintf(models.PushEventMsg, p.UserName, p.Project.PathWithNamespace, s[1], p.Project.Homepage, p.UserUsername, "no message commit", "no url")
	} else {
		res = fmt.Sprintf(models.PushEventMsg, p.UserName, p.Project.PathWithNamespace, s[1], p.Project.Homepage, p.UserUsername, p.Commits[0].Message, p.Commits[0].URL)
	}
	return msgResult{
		content: res,
		//url:     p.Commits[0].URL,
		//text:    "Open Commit",
		err: err,
	}
}
func issueMessage(body []byte) msgResult {
	var p webhook.Issues

	err := json.Unmarshal(body, &p)
	if err != nil {
		log.Panicln(err)
	}
	desc := strings.Split(p.ObjectAttributes.Description, "\n")
	var viewMore string
	var issueText string
	if len(desc) > 4 {
		viewMore = " ..."
		issueText = desc[0] + "\n" + desc[1] + "\n" + desc[2]
	} else {
		issueText = p.ObjectAttributes.Description
	}
	var actionissue string

	switch p.ObjectAttributes.Action {
	case "reopen":
		actionissue = "reopened"
	case "close":
		actionissue = "closed"
	case "update":
		var count int
		if p.Changes.Description.Previous != p.Changes.Description.Current && p.Changes.Description.Previous != nil {
			actionissue = "changed description of "
			count++

		}
		if p.Changes.Title.Previous != p.Changes.Title.Current && p.Changes.Title.Previous != nil {
			actionissue = "changed title into "
			count++

		}

		if count != 1 {
			if p.Changes.Description.Previous == "null" {
				actionissue = "created"
			} else {
				actionissue = "updated"
			}
		}
	case "open":
		actionissue = "created"
	}

	return msgResult{
		content: fmt.Sprintf(models.IssueEventMsg, p.ObjectAttributes.Iid, p.User.Name, actionissue, p.ObjectAttributes.Title, p.ObjectAttributes.URL, issueText, viewMore),
		//url:     p.ObjectAttributes.URL,
		//text:    "Open Issue",
	}
}
func commentMessage(body []byte) msgResult {
	var noteType comment.NoteableType
	var err error
	err = json.Unmarshal(body, &noteType)
	if err != nil {
		log.Panicln(err)
	}
	var content string

	switch noteType.ObjectAttributes.NoteableType {
	case "Commit":
		var p comment.Commit
		err = json.Unmarshal(body, &p)
		content = fmt.Sprintf(models.CmtCommitMsg, p.Commit.Author.Name, p.Commit.ID, p.Commit.URL, p.ObjectAttributes.Note)
		//url, text = p.Commit.URL, "Open Commit"
	case "Issue":
		var p comment.Issues
		err = json.Unmarshal(body, &p)
		content = fmt.Sprintf(models.CmtIssueMsg, p.User.Name, p.Issue.Iid, p.Issue.Title, p.ObjectAttributes.URL, p.ObjectAttributes.Note)
		//url, text = p.ObjectAttributes.URL, "Open Issue"
	case "MergeRequest":
		var p comment.MergeRequest
		err = json.Unmarshal(body, &p)
		statusMR := strings.Split(p.MergeRequest.Title, ":")
		draftMR := statusMR[0]
		if draftMR != "Draft" {
			content = fmt.Sprintf(models.CmtMergeMsg, p.User.Name, p.MergeRequest.Title, p.ObjectAttributes.URL, p.ObjectAttributes.Note)
			//url, text = p.ObjectAttributes.URL, "Open Merge Request"
		}
	case "Snippet":
		var p comment.CodeSnippet
		err = json.Unmarshal(body, &p)
		content = fmt.Sprintf(models.CmtSnippetMsg, p.User.Name, p.ObjectAttributes.ID, p.ObjectAttributes.URL, p.ObjectAttributes.Note)
		//url, text = p.ObjectAttributes.URL, "Open Code Snippet"
	default:
		log.Fatalln("Invalid noteable type.")
	}

	return msgResult{
		content: content,
		//text:    text,
		//url:     url,
		err: err,
	}
}
func mergeRequestMessage(body []byte) msgResult {
	var p webhook.MergeRequestEventsLoad
	err := json.Unmarshal(body, &p)
	statusMR := strings.Split(p.ObjectAttributes.Title, ":")
	draftMR := statusMR[0]
	var content string

	if p.Changes.Title.Current != p.Changes.Title.Previous {
		if draftMR != "Draft" {
			content = fmt.Sprintf(models.StatusReadyMrMsg, p.ObjectAttributes.SourceBranch, p.ObjectAttributes.URL)
			//url, text = p.ObjectAttributes.URL, "Open Request"
		} else {
			content = fmt.Sprintf(models.StatusDraftMrMsg, p.ObjectAttributes.SourceBranch, p.ObjectAttributes.URL)
			//url, text = p.ObjectAttributes.URL, "Open Request"
		}
	} else {
		if draftMR != "Draft" {
			content = fmt.Sprintf(models.MergeRequestEventsMsg, p.User.Username, p.ObjectAttributes.Action, p.User.Username, p.Project.Name, p.ObjectAttributes.SourceBranch, p.ObjectAttributes.URL)
			//url, text = p.ObjectAttributes.URL, "Open Request"
		}
	}

	return msgResult{
		content: content,
		err:     err,
	}
}
func wikiPageMessage(body []byte) msgResult {
	var p webhook.WikipageEventsLoad
	err := json.Unmarshal(body, &p)
	content := fmt.Sprintf(models.WikipageEventsMsg, p.User.Username, p.ObjectAttributes.Title)
	//url, text := p.ObjectAttributes.URL, "Open WikiPage"

	return msgResult{
		content: content,
		//text:    text,
		//url:     url,
		err: err,
	}
}
func tagPushMessage(body []byte) msgResult {
	var p webhook.TagEventsLoad
	err := json.Unmarshal(body, &p)
	content := fmt.Sprintf(models.TagEventsMsg, p.UserName, p.Ref)
	//url, text := p.Project.WebURL, "Open WikiTag"

	return msgResult{
		content: content,
		//text:    text,
		//url:     url,
		err: err,
	}
}

// func buildMessage(body []byte) msgResult {
// 	var p webhook.JobsEvent
// 	err := json.Unmarshal(body, &p)
// 	content := fmt.Sprintf(models.JobsEvent, buildEmoji[p.BuildStage], p.BuildName, p.Ref, statusEmoji[p.BuildStatus], p.BuildStatus)

//		return msgResult{
//			content: content,
//			err:     err,
//		}
//	}
func featFlagMessage(body []byte) msgResult {
	var p webhook.FeatureFlag
	err := json.Unmarshal(body, &p)
	var s string
	if p.ObjectAttributes.Active {
		s = "active"
	} else {
		s = "unactive"
	}
	content := fmt.Sprintf(models.FeatFlagMsg, p.ObjectAttributes.Name, s)
	//url, text := p.Project.Homepage, "Open Project"

	return msgResult{
		content: content,
		//text:    text,
		//url:     url,
		err: err,
	}
}

// func pipelineMessage(body []byte) msgResult {
// 	var p webhook.PipelineEventsLoad
// 	err := json.Unmarshal(body, &p)
// 	content := fmt.Sprintf(models.PipelineEventsMsg, p.User.Username, p.ObjectAttributes.Ref, p.User.Username, p.Project.Name, p.Project.DefaultBranch, p.ObjectAttributes.Status)

//		return msgResult{
//			content: content,
//			err:     err,
//		}
//	}
func deployMessage(body []byte) msgResult {
	var p webhook.DeploymentEventsLoad
	err := json.Unmarshal(body, &p)
	if err != nil {
		log.Panicln(err)
	}

	return msgResult{
		content: fmt.Sprintf(models.DeployEventsMsg, p.User.Username, p.User.Username, p.Project.Name, p.Project.DefaultBranch, p.Project.Homepage, p.Status),
		//url:     p.Project.WebURL,
		//text:    "Open Deloyment",
		err: err,
	}
}
func releaseMessage(body []byte) msgResult {
	var p webhook.ReleaseEventsLoad
	err := json.Unmarshal(body, &p)
	if err != nil {
		log.Panicln(err)
	}

	return msgResult{
		content: fmt.Sprintf(models.ReleaseEventsMsg, p.Commit.Author.Name, p.Commit.Author.Name, p.Project.Name, p.Project.DefaultBranch, p.Project.Homepage, p.Commit.Message),
		//url:     p.Project.WebURL,
		//text:    "Open Release",
		err: err,
	}
}

func buildAction(message tgbotapi.Message, body []byte) string {
	var p webhook.JobsEvent
	//var content string
	if err := json.Unmarshal(body, &p); err != nil {
		panic(err)
	}

	return fmt.Sprintf(models.JobsEvent, buildEmoji[p.BuildStage], statusEmoji[p.BuildStatus], p.BuildName, p.BuildStatus)

}

// func pipelineAction(message tgbotapi.Message, body []byte) string {
// 	var p webhook.PipelineEventsLoad
// 	if err := json.Unmarshal(body, &p); err != nil {
// 		panic(err)
// 	}

// 	return fmt.Sprintf(models.PipelineEventsMsg, p.User.Username, p.ObjectAttributes.Ref, p.User.Username, p.Project.Name, p.Project.DefaultBranch, p.ObjectAttributes.Status)
// }

func SendMessage(chatID string, pay models.ObjectKind, body []byte) {
	id := atoi64(chatID)
	bot, err := LoadBot()
	if err != nil {
		panic(err)
	}

	res := func() (res msgResult) {
		if action, ok := sendActions[pay.ObjectKind]; ok {
			res = action(body)
		}
		return
	}()

	msg := tgbotapi.NewMessage(id, res.content)
	msg.ParseMode = "markdown"

	if res.text != "" || res.url != "" {
		msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
			InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
				{
					tgbotapi.InlineKeyboardButton{
						Text: res.text,
						URL:  &res.url,
					},
				},
			},
		}
	}

	message, err := bot.Send(msg)
	currentMesID = message.MessageID

	if pay.ObjectKind == "push" {
		configs.SetPushMessage(message, res.content)
	}

	if err != nil {
		panic(err)
	}
}
func EditMessage(chatID string, pay models.ObjectKind, body []byte) {

	id := atoi64(chatID)
	message := configs.GetPushMessage()

	if message.Text == "" && message.MessageID == 0 {
		return
	}
	bot, err := LoadBot()
	if err != nil {
		panic(err)
	}

	editText := func() (res string) {
		if action, ok := editActions[pay.ObjectKind]; ok {
			res = action(message, body)
		}
		return
	}()

	newText := message.Text + "\n\n" + editText
	msg := tgbotapi.NewEditMessageText(id, currentMesID, newText)
	msg.ParseMode = "markdown"
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}

}
func atoi64(s string) (i64 int64) {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panicln(err)
	}
	i64 = int64(i)
	return
}
