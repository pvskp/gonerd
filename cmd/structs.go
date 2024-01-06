package cmd

import "time"

type GitHubRepoResponse struct {
	Payload struct {
		AllShortcutsEnabled bool   `json:"allShortcutsEnabled"`
		Path                string `json:"path"`
		Repo                struct {
			ID                 int       `json:"id"`
			DefaultBranch      string    `json:"defaultBranch"`
			Name               string    `json:"name"`
			OwnerLogin         string    `json:"ownerLogin"`
			CurrentUserCanPush bool      `json:"currentUserCanPush"`
			IsFork             bool      `json:"isFork"`
			IsEmpty            bool      `json:"isEmpty"`
			CreatedAt          time.Time `json:"createdAt"`
			OwnerAvatar        string    `json:"ownerAvatar"`
			Public             bool      `json:"public"`
			Private            bool      `json:"private"`
			IsOrgOwned         bool      `json:"isOrgOwned"`
		} `json:"repo"`
		CurrentUser any `json:"currentUser"`
		RefInfo     struct {
			Name         string `json:"name"`
			ListCacheKey string `json:"listCacheKey"`
			CanEdit      bool   `json:"canEdit"`
			RefType      string `json:"refType"`
			CurrentOid   string `json:"currentOid"`
		} `json:"refInfo"`
		Tree struct {
			Items []struct {
				Name        string `json:"name"`
				Path        string `json:"path"`
				ContentType string `json:"contentType"`
			} `json:"items"`
			TemplateDirectorySuggestionURL any  `json:"templateDirectorySuggestionUrl"`
			Readme                         any  `json:"readme"`
			TotalCount                     int  `json:"totalCount"`
			ShowBranchInfobar              bool `json:"showBranchInfobar"`
		} `json:"tree"`
		FileTree struct {
			NAMING_FAILED struct {
				Items []struct {
					Name        string `json:"name"`
					Path        string `json:"path"`
					ContentType string `json:"contentType"`
				} `json:"items"`
				TotalCount int `json:"totalCount"`
			} `json:""`
		} `json:"fileTree"`
		FileTreeProcessingTime float64 `json:"fileTreeProcessingTime"`
		FoldersToFetch         []any   `json:"foldersToFetch"`
		TreeExpanded           bool    `json:"treeExpanded"`
		SymbolsExpanded        bool    `json:"symbolsExpanded"`
		CsrfTokens             struct {
			RyanoasisNerdFontsBranches struct {
				Post string `json:"post"`
			} `json:"/ryanoasis/nerd-fonts/branches"`
			RyanoasisNerdFontsBranchesFetchAndMergeMaster struct {
				Post string `json:"post"`
			} `json:"/ryanoasis/nerd-fonts/branches/fetch_and_merge/master"`
			RyanoasisNerdFontsBranchesFetchAndMergeMasterDiscardChangesTrue struct {
				Post string `json:"post"`
			} `json:"/ryanoasis/nerd-fonts/branches/fetch_and_merge/master?discard_changes=true"`
		} `json:"csrf_tokens"`
	} `json:"payload"`
	Title      string `json:"title"`
	AppPayload struct {
		HelpURL              string `json:"helpUrl"`
		FindFileWorkerPath   string `json:"findFileWorkerPath"`
		FindInFileWorkerPath string `json:"findInFileWorkerPath"`
		GithubDevURL         any    `json:"githubDevUrl"`
		EnabledFeatures      struct {
			CodeNavUIEvents                        bool `json:"code_nav_ui_events"`
			CopilotConversationalUx                bool `json:"copilot_conversational_ux"`
			CopilotConversationalUxEmbeddingUpdate bool `json:"copilot_conversational_ux_embedding_update"`
			CopilotConversationalUxStreaming       bool `json:"copilot_conversational_ux_streaming"`
			CopilotPopoverFileEditorHeader         bool `json:"copilot_popover_file_editor_header"`
			CopilotSmellIcebreakerUx               bool `json:"copilot_smell_icebreaker_ux"`
			ReactBlobSnakeSymbols                  bool `json:"react_blob_snake_symbols"`
		} `json:"enabled_features"`
	} `json:"appPayload"`
}
