local client = vim.lsp.start_client {
    name = "LSP",
    cmd = { "/home/kassab/GO/lsp/cmd/lsp" },
}

if not client then
    vim.notify("LSP client failed to start", vim.log.levels.ERROR)
    return
end

vim.api.nvim_create_autocmd("FileType", {
    pattern = "markdown",
    callback = function()
        vim.lsp.buf_attach_client(0, client)
    end

})