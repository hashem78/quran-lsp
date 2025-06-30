vim.api.nvim_create_autocmd({ "VimEnter", "BufReadPre", "FileType" }, {
	pattern = { "*.txt" },
	callback = function()
		local client_id = vim.lsp.start({
			name = "quran-lsp",
			cmd = { "/home/mythi/development/quran-lsp/language-server/quran-lsp" },
			filetypes = { "txt" },
			on_attach = require("nvim-plugin.lua.on-attach").on_lsp_attach,
		})
		if not client_id then
			vim.notify("Did not attach quarn-lsp")
			return
		end
		vim.lsp.buf_attach_client(0, client_id)
	end,
})
