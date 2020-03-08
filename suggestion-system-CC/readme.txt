This system consists of but is not minimized or maximized to 5 different custom commands

deletesuggestion-CC - deletes the suggestion that suggest-CC made
	Trigger type: command
	Trigger: deletesuggestion
	Command Usage: -deletesuggestion MESSAGE-ID

suggest-CC - creates a suggestion in provided channel
	Trigger type: command
	Trigger: suggest
	Command Usage: -suggest <Suggestion-Here>

suggestresult-CC - finds the ending result of the suggestion using the reaction count on the suggestion message
	Trigger type: command
	Trigger: result
	Command Usage: -result MESSAGE-ID
		       -result MESSAGE-ID <Outcome>
		       -result MessageID choice votes majority neutralvotes upvotes downvotes

suggestionimplement-CC - marks the suggestion as implemented
	Trigger type: command
	Trigger: implement
	Command Usage: -implement MESSAGE-ID

suggestiontracking-CC - tracks the count of votes on a suggestion starting at the number of positive votes you provided
	Trigger type: Reaction
	Trigger: added+removed reactions
	Command Usage: Get 10 people to react to a suggestion

Please make an issue or contact 𝕊𝕙𝕒𝕕𝕠𝕨ℕ𝕚𝕘𝕙𝕥#2845 (580567048745385994) if there are any questions, concerns, or suggestions.

