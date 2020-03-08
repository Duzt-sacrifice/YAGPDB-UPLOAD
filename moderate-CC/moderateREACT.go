{{/* Trigger type: Reacton  Trigger: Reactions added+removed */}}

{{/* Variables */}}
{{$mod := ROLE-ID}}
{{$admin := ROLE-ID}}
{{$modreact := (split (joinStr "" "❗ 🔇 🔈 👢 🔨 ❌") " ")}}
{{$123x := (split (joinStr "" "1️⃣ 2️⃣ 3️⃣ ❌") " ")}}
{{$mute := (joinStr "" "1️⃣ - Duration: 30 Minutes\n2️⃣ - Duration: 1 Hour\n3️⃣ - Duration: 24 Hours\n❌ - Back ")}}
{{$ban := (joinStr "" "1️⃣ - Duration: 24 Hours\n2️⃣ - Duration: 1 week\n3️⃣ - Duration: Permanently\n❌ - Back ")}}
{{$modmenu := (joinStr "" "❗ - Warn\n🔇 - Mute\n🔈 - Unmute\n👢 - Kick\n🔨 - Ban\n❌ - Exit\n \\> React to choose what action you want to perform < ")}}


{{/* Don't Mess with this unless you know what you are doing */}} {{/* Reaction */}}

{{define "ts1"}}
{{editMessage .Reaction.ChannelID (.Get "messageID") (.Get "newembed") }}
{{deleteAllMessageReactions .Reaction.ChannelID (.Get "messageID")}}
{{end}}
{{if or (hasRoleID $admin) (hasRoleID $mod)}}
	{{if .ReactionAdded}}
		{{if (dbGet .Reaction.UserID "modmenu").Value}}
			{{if eq (index (split ((dbGet .Reaction.UserID "modmenu").Value) ":") 0) (toString .Reaction.MessageID)}}
				{{$user := (userArg (index (split ((dbGet .Reaction.UserID "modmenu").Value) ":") 1)) }}
				{{$message := getMessage .Channel.ID (index (split ((dbGet .Reaction.UserID "modmenu").Value) ":") 0) }}
				{{$embed := index $message.Embeds 0}}
				{{if or (eq .Reaction.Emoji.Name "❗") (eq .Reaction.Emoji.Name "🔇") (eq .Reaction.Emoji.Name "🔈") (eq .Reaction.Emoji.Name "👢") (eq .Reaction.Emoji.Name "🔨")}}
					{{if and (eq .Reaction.Emoji.Name "❗") (eq $embed.Title "Moderation Menu")}}
						{{$silent := exec "warn" $user "Warned By Mod Menu"}}
						{{$newembed := cembed "title" "Moderation Menu" "description" $modmenu "color" 77}}
						{{template "ts1" (sdict "messageID" $message.ID "newembed" $newembed) }}
						{{addMessageReactions .Reaction.ChannelID $message.ID $modreact}}
						Warned User, Use `-reason Mod-Log-MessageID Reason-Specified-Here` to add a reason {{deleteResponse 10}}
					{{end}}
					{{if and (eq .Reaction.Emoji.Name "🔇") (eq $embed.Title "Moderation Menu")}}
						{{$newembed := cembed "title" "Mute Duration" "description" $mute "color" 77}}
						{{template "ts1" (sdict "messageID" $message.ID "newembed" $newembed) }}
						{{addMessageReactions .Reaction.ChannelID $message.ID $123x}}
					{{end}}
					{{if and (eq .Reaction.Emoji.Name "🔈") (eq $embed.Title "Moderation Menu")}}
						{{$silent := exec "unmute" $user "Unmuted By Mod Menu"}}
						{{$newembed := cembed "title" "Moderation Menu" "description" $modmenu "color" 77}}
						{{template "ts1" (sdict "messageID" $message.ID "newembed" $newembed) }}
						{{addMessageReactions .Reaction.ChannelID $message.ID $modreact}}
						Unmuted User, Use `-reason Mod-Log-MessageID Reason-Specified-Here` to add a reason {{deleteResponse 10}}
					{{end}}
					{{if and (eq .Reaction.Emoji.Name "👢") (eq $embed.Title "Moderation Menu")}}
						{{$silent := exec "kick" $user "Kicked By Mod Menu"}}
						{{$newembed := cembed "title" "Moderation Menu" "description" $modmenu "color" 77}}
						{{template "ts1" (sdict "messageID" $message.ID "newembed" $newembed) }}
						{{addMessageReactions .Reaction.ChannelID $message.ID $modreact}}
						Kicked User, Use `-reason Mod-Log-MessageID Reason-Specified-Here` to add a reason {{deleteResponse 10}}
					{{end}}
					{{if and (eq .Reaction.Emoji.Name "🔨") (eq $embed.Title "Moderation Menu")}}
						{{$newembed := cembed "title" "Ban Duration" "description" $ban "color" 77}}
						{{template "ts1" (sdict "messageID" $message.ID "newembed" $newembed) }}
						{{addMessageReactions .Reaction.ChannelID $message.ID $123x}}
					{{end}}
				{{end}}
				{{if eq .Reaction.Emoji.Name "1️⃣"}}
					{{if eq $embed.Title "Mute Duration"}}
						{{$silent := exec "mute" $user.ID "30m" "Muted By Mod Menu"}}
						{{$newembed := cembed "title" "Moderation Menu" "description" $modmenu "color" 77}}
						{{template "ts1" (sdict "messageID" $message.ID "newembed" $newembed) }}
						{{addMessageReactions .Reaction.ChannelID $message.ID $modreact}}
						Muted User for 1 hour, Use `-reason Mod-Log-MessageID Reason-Specified-Here` to add a reason {{deleteResponse 10}}
					{{end}}
					{{if eq $embed.Title "Ban Duration"}}
						{{$silent := exec "ban" $user.ID "Banned By Mod Menu" "-d 24h"}}
						{{$newembed := cembed "title" "Moderation Menu" "description" $modmenu "color" 77}}
						{{template "ts1" (sdict "messageID" $message.ID "newembed" $newembed) }}
						{{addMessageReactions .Reaction.ChannelID $message.ID $modreact}}
						Banned User for 1 hour, Use `-reason Mod-Log-MessageID Reason-Specified-Here` to add a reason {{deleteResponse 10}}
					{{end}}
				{{end}}
				{{if eq .Reaction.Emoji.Name "2️⃣"}}
					{{if eq $embed.Title "Mute Duration"}}
						{{$silent := exec "mute" $user.ID "1h" "Muted By Mod Menu"}}
						{{$newembed := cembed "title" "Moderation Menu" "description" $modmenu "color" 77}}
						{{template "ts1" (sdict "messageID" $message.ID "newembed" $newembed) }}
						{{addMessageReactions .Reaction.ChannelID $message.ID $modreact}}
						Muted User for 2 hours, Use `-reason Mod-Log-MessageID Reason-Specified-Here` to add a reason {{deleteResponse 10}}
					{{end}}
					{{if eq $embed.Title "Ban Duration"}}
						{{$silent := exec "ban" $user.ID "Banned By Mod Menu" "-d 1w"}}
						{{$newembed := cembed "title" "Moderation Menu" "description" $modmenu "color" 77}}
						{{template "ts1" (sdict "messageID" $message.ID "newembed" $newembed) }}
						{{addMessageReactions .Reaction.ChannelID $message.ID $modreact}}
						Banned User for 2 hours, Use `-reason Mod-Log-MessageID Reason-Specified-Here` to add a reason {{deleteResponse 10}}
					{{end}}
				{{end}}
				{{if eq .Reaction.Emoji.Name "3️⃣"}}
					{{if eq $embed.Title "Mute Duration"}}
						{{$silent := exec "mute" $user.ID "24h" "Muted By Mod Menu"}}
						{{$newembed := cembed "title" "Moderation Menu" "description" $modmenu "color" 77}}
						{{template "ts1" (sdict "messageID" $message.ID "newembed" $newembed) }}
						{{addMessageReactions .Reaction.ChannelID $message.ID $modreact}}
						Muted User for 3 hours, Use `-reason Mod-Log-MessageID Reason-Specified-Here` to add a reason {{deleteResponse 10}}
					{{end}}
					{{if eq $embed.Title "Ban Duration"}}
						{{$silent := exec "ban" $user.ID "Banned By Mod Menu" "-d p"}}
						{{$newembed := cembed "title" "Moderation Menu" "description" $modmenu "color" 77}}
						{{template "ts1" (sdict "messageID" $message.ID "newembed" $newembed) }}
						{{addMessageReactions .Reaction.ChannelID $message.ID $modreact}}
						Banned User for 3 hours, Use `-reason Mod-Log-MessageID Reason-Specified-Here` to add a reason {{deleteResponse 10}}
					{{end}}
				{{end}}
				{{if eq .Reaction.Emoji.Name "❌"}}
					{{if or (eq $embed.Title "Mute Duration") (eq $embed.Title "Ban Duration")}}
						{{$newembed := cembed "title" "Moderation Menu" "description" $modmenu "color" 77}}
						{{template "ts1" (sdict "messageID" $message.ID "newembed" $newembed) }}
						{{addMessageReactions .Reaction.ChannelID $message.ID $modreact}}
					{{end}}
					{{if eq $embed.Title "Moderation Menu"}}
						{{deleteMessage .Reaction.ChannelID $message.ID 1}}
						{{dbDel .Reaction.UserID "modmenu"}}
					{{end}}
				{{end}}
			{{end}}
		{{end}}
	{{end}}
{{end}}