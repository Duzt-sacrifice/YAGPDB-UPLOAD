{{/* Initializing Variable */}}

{{/* Amount of reaction needed */}}
{{$count := NUMBER}}


{{/* Two emote we are tracking in this scenario, you may need to change the code if you are not tracking two emotes*/}}
{{$emoteNameUp := "✅"}}
{{$emoteNameNeutral := "neutralvote"}}
{{$emoteNameDown:= "downvote"}}

{{/* Confirmation Emote, emote to inform that the msg has been pinned, leave ID blank if default emote*/}}
{{$emoteNameAdded:= "⬆️"}}
{{$emoteIdAdded := ""}}

{{/* Channel we repost msg in with upvote and downvote count */}}
{{$trackingchannel := CHANNEL-ID}}



{{/* DO NOT TOUCH BELOW UNLESS YOU KNOW WHAT YOU ARE DOING */}}


{{/* INIT VARIABLE */}}
{{$countUpvote := 0}}
{{$countNeutral := 0}}
{{$countDownvote := 0}}
{{$match := 0}}



{{/* ITERATING THROUGH REACTION TO CHECK IF MEETS CONDITION AND UPDATE VARIABLES*/}}


{{range .ReactionMessage.Reactions}}


{{/* If emote names matches the emote we are tracking */}}

{{if eq (toString .Emoji.Name) (toString $emoteNameDown)}}
{{/* Update count with the count of the emote */}}
{{$countDownvote = (sub .Count 1)}}
{{end}}

{{if eq (toString .Emoji.Name) (toString $emoteNameNeutral)}}
{{$countNeutral = (sub .Count 1)}}
{{end}}

{{if eq (toString .Emoji.Name) (toString $emoteNameUp)}}
{{/* if suggestion emote meets required count, its a match and we proceed pass the loop*/}}
{{if gt .Count $count}}
{{$countUpvote = (sub .Count 1)}}
{{$match = 1}}
{{end}}{{end}}{{end}}

{{/* Match tell us the suggestion emote met the required count and a message should be updated or created.  */}}
{{if (eq $match 1) }}


{{/* CREATING THE MESSAGE*/}}


{{$tmpString := ""}}

{{$tmpString = (printf "**[Message Link](https://discordapp.com/channels/%d/%d/%d) %d Upvote %d Neutral %d Downvote**\n"  .ReactionMessage.GuildID .ReactionMessage.ChannelID  .ReactionMessage.ID $countUpvote $countNeutral $countDownvote )}}


{{$RawEmbed := sdict}}

{{/* If the message is a Embed we transpose the embed into a new one with the message link and count*/}}

{{if .ReactionMessage.Embeds}}
{{$embed := (index .ReactionMessage.Embeds 0)}}

{{if $embed.Title}}
{{$RawEmbed.Set "title" ($embed.Title)}}
{{end}}

{{if $embed.Description}}
{{$RawEmbed.Set "description" (joinStr "" $tmpString ($embed.Description))}}
{{else}}
{{$RawEmbed.Set "description" ($tmpString)}}
{{end}}

{{if $embed.Thumbnail}}
{{$RawEmbed.Set "thumbnail" ($embed.Thumbnail)}}
{{end}}

{{if $embed.Fields}}
{{$RawEmbed.Set "fields" ($embed.Fields)}}
{{end}}

{{if $embed.URL}}
{{$RawEmbed.Set "url" ($embed.URL)}}
{{end}}



{{/* If message is not a embed, we copy the information needed such as author and content. Later we add emote count and message link */}}

{{if $embed.Author}}
{{$RawEmbed.Set "author" (sdict "name" (toString $embed.Author.Name) "url" (toString ($embed.Author.URL) ) "icon_url" (toString ($embed.Author.IconURL))  )}}
{{else}}
{{$RawEmbed.Set "author" (sdict "name" "")}}
{{end}}

{{/* Handling Embeds with Message Content */}}
{{if .ReactionMessage.Content}}
{{$temp:= $RawEmbed.Get "description"}}
{{$RawEmbed.Set "description" (joinStr "" $tmpString $temp "\n```Message Content```\n" .ReactionMessage.Content )}}
{{end}}

{{end}}



{{/* CHECKING IF CREATING A NEW MESSAGE OR EDITING EXISTING MESSAGE */}}


{{/* looping again through all the reaction to check for the added confirmation emote */}}
{{$dup := 0}}
{{range .ReactionMessage.Reactions}}
{{if eq (toString .Emoji.Name) (toString $emoteNameAdded)}}
{{if not .Me}}
{{else}}

{{/* Found the emote and confirmed that the bot reacted to it, not some random person*/}}
{{$dup = 1}}
{{end}}
{{end}}
{{end}}
{{if eq $dup 0}}


{{/* No dup found, sending the embed to the appropriate channel and react to show completion */}}
{{$msgID := sendMessageRetID (toString $trackingchannel) (cembed $RawEmbed) }}

{{/* Update Db with the embed message ID for when we need to edit and update the emote numbers */}}
{{dbSet .Guild.ID (.Message.ID) (toString $msgID)}}

{{/* adding confirmation emote */}}
{{/* If its a default emote, there is no ID */}}
{{if eq (len $emoteIdAdded) 0}}
{{ addReactions $emoteNameAdded}}
{{else}}
{{$fullEmote := (joinStr "" $emoteNameAdded ":" $emoteIdAdded)}}
{{ addReactions $fullEmote}}
{{end}}
{{else}}


{{/* Message already exist, don't send new message, update the existing one, */}}

{{/* Get messageID from Database and editmessage */}}
{{$editMessageID := dbGet .Guild.ID .Message.ID}}
{{editMessage (toString $trackingchannel) (toString $editMessageID.Value)  (cembed $RawEmbed)}}

{{end}}
{{end}}