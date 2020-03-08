{{/* Trigger type: command  Trigger: message 
	
	This fetches the text and/or embed that is on the message you provide, Command usage: -message channel-id message-id */}}

{{$args := parseArgs 2 ""
	(carg "int" "channel provided")
	(carg "int" "message provided")}}

{{$cp := ($args.Get 0)}}
{{$mp := ($args.Get 1)}}

{{/* message building */}}
{{$content := ""}}
{{$embed := ""}}

{{if and ((getMessage $cp $mp).Content) (index (getMessage $cp $mp).Embeds)}}
    {{$content = (getMessage $cp $mp).Content}}
    {{$embed = (index (getMessage $cp $mp).Embeds 0)}}
    {{sendMessage nil (complexMessage "content" (joinStr "" "<#" $cp "> : " (getMessage $cp $mp).Author.String " (" (getMessage $cp $mp).Author.ID ") : " ((getMessage $cp $mp).Timestamp.Parse.Format "1/02/2006 3:04PM") " UTC" "\n\n" ($content)) "embed" ($embed))}}
{{else}}
    {{if (getMessage $cp $mp).Content}}
        {{$content = (getMessage $cp $mp).Content}}
        {{sendMessage nil (joinStr "" "<#" $cp "> : " (getMessage $cp $mp).Author.String " (" (getMessage $cp $mp).Author.ID ") : " ((getMessage $cp $mp).Timestamp.Parse.Format "1/02/2006 3:04PM") " UTC" "\n\n" ($content))}}
    {{end}}
    {{if (index (getMessage $cp $mp).Embeds)}}
        {{$embed = (index (getMessage $cp $mp).Embeds 0)}}
        {{sendMessage nil (complexMessage "content" (joinStr "" "<#" $cp "> : " (getMessage $cp $mp).Author.String " (" (getMessage $cp $mp).Author.ID ") : " ((getMessage $cp $mp).Timestamp.Parse.Format "1/02/2006 3:04PM") " UTC") "embed" ($embed))}}
    {{end}}
{{end}}