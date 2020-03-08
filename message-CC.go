{{/* Trigger type: command  Trigger: message 
	
	This fetches the text and/or embed that is on the message you provide, Command usage: -message channel-id message-id */}}

{{$args := parseArgs 2 ""
	(carg "int" "channel provided")
    (carg "int" "message provided")}}

{{$cp := ($args.Get 0)}}
{{$mp := ($args.Get 1)}}

{{/* message building */}}
{{$content := "There is no content on this message"}}
{{$embed := (cembed "description" "There is no embed on this message")}}

{{if (getMessage $cp $mp).Content}}
    {{$content = (getMessage $cp $mp).Content}}
{{end}}
{{if (index (getMessage $cp $mp).Embeds 0)}}
    {{$embed = (index (getMessage $cp $mp).Embeds 0)}}
{{end}}

{{sendMessage nil (complexMessage "content" ($content) "embed" ($embed))}}