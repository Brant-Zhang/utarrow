""""""""""""""""""""""
"     plug
"
""""""""""""""""""""
call plug#begin()
Plug 'fatih/vim-go', { 'do': ':GoInstallBinaries' }
Plug 'AndrewRadev/splitjoin.vim'
"Plug 'SirVer/ultisnips'
Plug 'ctrlpvim/ctrlp.vim'
Plug 'Tagbar'
Plug 'scrooloose/nerdtree'
Plug 'https://github.com/Valloric/YouCompleteMe.git'
"Plug 'godlygeek/tabular'
"Plug 'plasticboy/vim-markdown'
call plug#end()


"""""""""""""""""""""
"      base setting
""""""""""""""""""""""
set autochdir
"set backspace=2
"设置vim粘贴和系统剪切板公用
set clipboard+=unnamed
set nocompatible                " Enables us Vim specific features
filetype off                    " Reset filetype detection first ...
filetype plugin indent on       " ... and enable filetype detection
set encoding=utf-8              " Set default encoding to UTF-8
set autoread                    " Automatically read changed files
set autoindent                  " Enabile Autoindent
set backspace=indent,eol,start  " Makes backspace key more powerful.
set ttyfast                     " Indicate fast terminal conn for faster redraw
set ttymouse=xterm2             " Indicate terminal type for mouse codes
set ttyscroll=3                 " Speedup scrolling
set incsearch                   " Shows the match while typing
set hlsearch                    " Highlight found searches
set noerrorbells                " No beeps
set number                      " Show line numbers
set showcmd                     " Show me what I'm typing
set noswapfile                  " Don't use swapfile
set nobackup                    " Don't create annoying backup files
set autowrite                   " Automatically save before :next, :make etc.
set hidden                      " Buffer should still exist if window is closed
set fileformats=unix,dos,mac    " Prefer Unix over Windows over OS 9 formats
set noshowmatch                 " Do not show matching brackets by flickering
set noshowmode                  " We show the mode with airline or lightline
set ignorecase                  " Search case insensitive...
set smartcase                   " ... but not it begins with upper case
set completeopt=menu,menuone    " Show popup menu, even if there is one entry
set pumheight=10                " Completion window max size
set nocursorcolumn              " Do not highlight column (speeds up highlighting)
set nocursorline                " Do not highlight cursor (speeds up highlighting)
set lazyredraw                  " Wait to redraw

" Colorscheme
"syntax enable
"set t_Co=256
"let g:rehash256 = 1
"let g:molokai_original = 1
"colorscheme molokai

" Visual linewise up and down by default (and use gj gk to go quicker)
"noremap <Up> gk
"noremap <Down> gj
"noremap j gj
"noremap k gk

"""""""""""""""""""""""""""""""
"         go
""""""""""""""""""""""""""""""
let g:go_highlight_types = 1
let g:go_highlight_functions = 1
let g:go_highlight_methods = 1
let g:go_highlight_fields = 1
let g:go_highlight_operators = 1
let g:go_highlight_build_constraints = 1
autocmd BufNewFile,BufRead *.go setlocal noexpandtab tabstop=4 shiftwidth=4 
let g:go_fmt_command = "goimports"
" set mapleader
let mapleader = ","


"autocmd vimenter * NERDTree
"----------------------------------------

"set tagbar windows kuaidu
let g:tagbar_width=30
let g:tagbar_type_go = {
	\ 'ctagstype' : 'go',
	\ 'kinds'     : [
		\ 'p:package',
		\ 'i:imports:1',
		\ 'c:constants',
		\ 'v:variables',
		\ 't:types',
		\ 'n:interfaces',
		\ 'w:fields',
		\ 'e:embedded',
		\ 'm:methods',
		\ 'r:constructor',
		\ 'f:functions'
	\ ],
	\ 'sro' : '.',
	\ 'kind2scope' : {
		\ 't' : 'ctype',
		\ 'n' : 'ntype'
	\ },
	\ 'scope2kind' : {
		\ 'ctype' : 't',
		\ 'ntype' : 'n'
	\ },
	\ 'ctagsbin'  : 'gotags',
	\ 'ctagsargs' : '-sort -silent'
	\ }

map <F8> :TagbarToggle<CR>
map <F7> :NERDTreeToggle<CR>

"---------------------------------------------------------------
" YCM settings
let g:ycm_key_invoke_completion = '<C-Space>'
let g:ycm_key_list_select_completion = ['', '']
let g:ycm_key_list_previous_completion = ['']

"---------------------------------------------------------------
" UltiSnips settings
let g:UltiSnipsExpandTrigger="<tab>"
let g:UltiSnipsJumpForwardTrigger="<c-k>"
let g:UltiSnipsJumpBackwardTrigger="<c-j>" 

