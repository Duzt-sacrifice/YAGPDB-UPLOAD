{{$suggestionchannel := CHANNEL-ID}} {{/* change this one line */}}

{{$args := parseArgs 1 "`-suggest <Suggestion:Text>`"
	(carg "string" "suggestion")}}

{{ $embed := cembed
"description" (joinStr " " .CmdArgs)
"color" 77
"author" (sdict "name" (joinStr "" .User.Username "#" .User.Discriminator) "url" "" "icon_url" (.User.AvatarURL "512"))
"timestamp"  currentTime
"footer" (sdict "text" (joinStr "" "Submit your suggestion with -suggest | " .User.ID))
}}
{{ $id := (sendMessageRetID $suggestionchannel $embed) }}
{{ addMessageReactions $suggestionchannel $id "✅" "neutralvote:665908078834614312" "downvote:665908071381336084" }}
{{ sendDM (joinStr "" "Suggestion submitted successfully. If you want to discuss this or other suggestions, use the <#634934110409850901> channel. If you want to delete your suggestion, do so with `-deletesuggestion " $id "` in the " .Guild.Name " server.") }}
{{addReactions "✅" }} {{deleteTrigger 20}}