{{deleteTrigger 1}}{{deleteResponse 3}}
{{if .CmdArgs}} 

{{/* dont change anything unless you know what you are doing */}}

	{{$suggestionsChannel := CHANNEL-ID}} {{/* change this */}}
	
{{/* dont change anything unless you know what you are doing */}}

    {{$message := getMessage $suggestionsChannel (index .CmdArgs 0)}}
    {{$embed := index $message.Embeds 0}}
    {{if $message}}
        {{$color := ""}}
        {{$choice := ""}}
        {{$Majority := ""}}
        {{$countNeutral := ""}}
        {{$countUpvote := ""}}
        {{$countDownvote := ""}}
        {{$Votes := ""}}
        {{if eq (len .Args) 8}}
            {{$choice = (title (lower (index .Args 2)))}} {{/*choice votes majority neutralvotes upvotes downvotes */}}
            {{$Votes = (index .Args 3)}}
            {{$Majority = (index .Args 4)}}
            {{$countNeutral = (index .Args 5)}}
            {{$countUpvote = (index .Args 6)}}
            {{$countDownvote = (index .Args 7)}}
            {{if or (eq $choice "Accepted") (eq $choice "Denied") (eq $choice "Undecided")}}    
                {{if eq $choice "Denied"}}
                    {{$color = 16212532}}
                {{else if eq $choice "Accepted"}}
                    {{$color = 6093439}}
                {{else if eq $choice "Undecided"}}
					{{$color = 16110134}}
                {{end}}
            {{else}}
                Command Usage: `-result MessageID choice votes majority neutralvotes upvotes downvotes`
            {{end}}
        {{else}}
            {{if le (len .CmdArgs) 3}}
                {{$emoteNameUp := "✅"}}
                {{$emoteNameNeutral := "neutralvote"}}
                {{$emoteNameDown:= "downvote"}}
                {{$countUpvote = 0}}
                {{$countNeutral = 0}}
                {{$countDownvote = 0}}
                {{range $message.Reactions}}
                    {{if eq (toString .Emoji.Name) (toString $emoteNameDown)}}
                        {{$countDownvote = (sub .Count 1)}}
                    {{end}}
                    {{if eq (toString .Emoji.Name) (toString $emoteNameNeutral)}}
                        {{$countNeutral = (sub .Count 1)}}
                    {{end}}
                    {{if eq (toString .Emoji.Name) (toString $emoteNameUp)}}
                        {{$countUpvote = (sub .Count 1)}}
                    {{end}}
                {{end}}
                {{if eq (len .Args) 2}}
                    {{$count := (sub $countUpvote $countDownvote)}}
                    {{if gt $count 0}}
                        {{$color = 6093439}}
                        {{$Majority = (toString $countUpvote)}}
                        {{$choice = "Accepted"}}
                    {{else if lt $count 0}}
                        {{$color = 16212532}}
                        {{$Majority = (toString $countDownvote)}}
                        {{$choice = "Denied"}}
                    {{else if eq $count 0}}
                        {{$color = 16110134}}
                        {{$Majority = (toString $countNeutral)}}
                        {{$choice = "Undecided"}}
                    {{end}}
                    {{$Votes = (toString (add $countUpvote $countNeutral $countDownvote))}}
                {{else if eq (len .Args) 3}}
                    {{$choice = (title (lower (index .CmdArgs 1)))}}
                    {{if or (eq $choice "Accepted") (eq $choice "Denied") (eq $choice "Undecided")}}    
                        {{if eq $choice "Denied"}}
                            {{$color = 16212532}}
                        {{else if eq $choice "Accepted"}}
                            {{$color = 6093439}}
                        {{else if eq $choice "Undecided"}}
							{{$color = 16110134}}
                        {{end}}
                        {{$count := (sub $countUpvote $countDownvote)}}
                        {{if gt $count 0}}
                            {{$Majority = (toString $countUpvote)}}
                        {{else if lt $count 0}}
                            {{$Majority = (toString $countDownvote)}}
                        {{else if eq $count 0}}
                            {{$Majority = (toString $countNeutral)}}
                        {{end}}
                        {{$Votes = (toString (add $countUpvote $countNeutral $countDownvote))}}
                    {{end}}
                {{end}}
            {{else}}
                error
            {{end}}
        {{end}}
        {{$newEmbed := cembed 
        "author" $embed.Author 
        "description" $embed.Description 
        "footer" $embed.Footer 
        "color" (toInt $color)
        "fields" (cslice 
            (sdict "name" "Voters" "value" (toString $Votes) "inline" true)
            (sdict "name" "Favourable" "value" (toString $countUpvote) "inline" true)
            (sdict "name" "Abstaineds" "value" (toString $countNeutral) "inline" true)
            (sdict "name" "Opposed" "value" (toString $countDownvote) "inline" true)
            (sdict "name" "Majority" "value" (toString $Majority) "inline" true)
            (sdict "name" "Outcome" "value" (toString $choice) "inline" true))}}
        {{editMessage $suggestionsChannel $message.ID $newEmbed}}
        {{deleteAllMessageReactions $suggestionsChannel $message.ID}}
        {{sleep 3}}
        Done
    {{end}}
{{end}}