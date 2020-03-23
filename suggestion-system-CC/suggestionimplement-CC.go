{{$suggestionsChannel := CHANNEL-ID}}
{{$archiveChannel := CHANNEL-ID}} 

{{/* Change the above only (unless you know what you are doing) */}}

{{$message :=  (getMessage $suggestionsChannel (index  .CmdArgs 0))}}
{{$embed := index $message.Embeds 0}}
{{$args := (joinStr " " .CmdArgs)}}

{{deleteTrigger 1}}{{deleteResponse 3}}
{{if (index .CmdArgs 0)}}{{$message := getMessage $suggestionsChannel (index .CmdArgs 0)}}
{{if $message}}
{{deleteMessage $suggestionsChannel (index .CmdArgs 0)}}{{sleep 3}}
{{$embed := cembed "title" "Successfully Implemented Suggestion:" "description" (index $message.Embeds 0).Description "author" (sdict "name" (index $message.Embeds 0).Author.Name) "footer" (sdict "text" (joinStr "" "Implemented by: " .User.Username " - " .User.ID))
"fields" $embed.Fields
"timestamp" $message.Timestamp}}
{{sendMessage $archiveChannel $embed}}
done
{{else}}No message ID provided!{{end}}
{{else}}Unknown message{{end}}
