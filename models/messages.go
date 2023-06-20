package models

var (
	PushEventMsg          = "⬆️ *Push Event*\n%s pushed to [%s/%s](%s) \n%s : [%s](%s)"
	IssueEventMsg         = "🚨 *Issue* %d \n %s %s [%s](%s).\n*Description*\n%s %s"
	MergeRequestEventsMsg = "🔀 *Merge Request*\n*%s* %s in [%s/%s/%s](%s) \n* a merge request*."
	// PipelineEventsMsg     = "▶️ *Pipeline*\n%s pipeline in branch %s in [%s/%s](%s) \n%s."
	DeployEventsMsg   = "🚀 *Deploy*\n%s deploy production in [%s/%s/%s](%s) \n %s."
	ReleaseEventsMsg  = "💿 *Release*\n%s release project in [%s/%s/%s](%s) \n*message*: `%s`."
	WikipageEventsMsg = "📝 *WikiPage*\n%s created a WikiPage with title %s."
	TagEventsMsg      = "🔖 *Tag*\n%s created a tag in branch %s."
	CmtCommitMsg      = "✍️ *Comment on the commit*\n%s commented on the commit [%s](%s).\n\n%s"
	CmtIssueMsg       = "✍️ *Comment at an issue*\n%s commented at an issue: [#%d %s](%s).\n\n%s"
	CmtMergeMsg       = "✍️ *Comment on a merge request*\n%s commented on a merge request [%s](%s).\n\n%s"
	CmtSnippetMsg     = "✍️ *Comment on a snippet*\n%s commented on a code snippet [#%s](%s).\n\n%s"
	JobsEvent         = "| %s%s %s is %s."
	FeatFlagMsg       = "🚩 *Feature Flag*\n%s is %s."
	StatusDraftMrMsg  = "🗒️ Merge Request [%s](%s) is marked as Draft "
	StatusReadyMrMsg  = "✅ Merge Request [%s](%s) is marked as Ready"
)

var (
	ChatExistMsg  = "💡 *Your chat is already added!*\n👉 Use this Webhook URL above to setup the notification.\n\n`https://%s/%s/%s`"
	ChatInsertMsg = "🎉 *Your chat is ready to receive Gitlab notification!*\n👉 To setup notifications *for this chat* with your GitLab repository, open Settings/Webhooks and add this URL:\n\n`https://%s/%s/%s`"
	ChatNotCmdMsg = "🥴 *Invalid Command.*\n👉 To interact with me:\n\n🎉 `/start` to get a webhook link.\n💥 `/drop` to drop webhook link."
	ChatDropMsg   = "💥 *Your notification is dropped.*\n👉 you'll no longer receive any message until you start a new one."
)
