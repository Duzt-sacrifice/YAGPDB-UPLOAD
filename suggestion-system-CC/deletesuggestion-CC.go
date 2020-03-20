{{sleep 1}}{{$args := parseArgs 1 "" (carg "int" "message id")}}

{{$suggestionsChannel := CHANNEL-ID}} {{/* change this one line */}}

{{$message := getMessage $suggestionsChannel ($args.Get 0)}}
{{if $message}}
    {{if $message.Embeds}}
        {{if gt (len (index $message.Embeds 0).Footer.Text) 37}}
            {{$suggestionAuthorID := (slice (index $message.Embeds 0).Footer.Text 39)}}
            {{if eq $suggestionAuthorID (toString .User.ID)}}
                {{ deleteMessage $suggestionsChannel $message.ID 1 }}
                Suggestion deleted.
            {{else}}
                You can only delete your own suggestions.
            {{end}}
        {{else}}
            suggestion command is set up incorrectly ping/dm a staff member
        {{end}}
    {{else}}
        Please provide a valid message ID (with embed).
    {{end}}
{{else}}
    Please provide a valid message ID.
{{end}}
{{deleteTrigger 5}} {{deleteResponse 5}}