# VS Code Setup for NoMercyChain Development

This guide will help you set up Visual Studio Code for NoMercyChain development on Windows.

## Prerequisites

- Visual Studio Code installed
- Go 1.18 or higher installed
- Node.js 16 or higher installed
- Git installed

## Setting Up VS Code

### 1. Install Required Extensions

Open VS Code and install the following extensions:

1. **Go** by Go Team at Google
   - Press `Ctrl+Shift+X` to open the Extensions view
   - Search for "Go"
   - Install the extension by Go Team at Google

2. **ESLint** for JavaScript linting
   - Search for "ESLint"
   - Install the extension by Microsoft

3. **Prettier** for code formatting
   - Search for "Prettier"
   - Install the extension by Prettier

4. **React** for React development
   - Search for "React"
   - Install "ES7+ React/Redux/React-Native snippets" by dsznajder

5. **Material Icon Theme** (optional, for better file icons)
   - Search for "Material Icon Theme"
   - Install the extension by Philipp Kief

### 2. Configure Go Tools

1. Open the Command Palette (`Ctrl+Shift+P`)
2. Type "Go: Install/Update Tools"
3. Select all tools and click "OK"

This will install essential Go tools like:
- gopls (Go language server)
- dlv (debugger)
- goimports (import formatter)
- golint (linter)
- and more

### 3. Configure Workspace Settings

1. Create a `.vscode` folder in your project root if it doesn't exist
2. Create a `settings.json` file inside the `.vscode` folder with the following content:

```json
{
  "go.useLanguageServer": true,
  "go.lintTool": "golint",
  "go.formatTool": "goimports",
  "editor.formatOnSave": true,
  "editor.codeActionsOnSave": {
    "source.organizeImports": true
  },
  "javascript.format.enable": false,
  "typescript.format.enable": false,
  "eslint.validate": ["javascript", "javascriptreact"],
  "editor.defaultFormatter": null,
  "[javascript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[javascriptreact]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[json]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[go]": {
    "editor.defaultFormatter": "golang.go"
  }
}
```

### 4. Create Launch Configurations

Create a `launch.json` file inside the `.vscode` folder with the following content:

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Go Backend",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/cmd/nmxchaind",
      "args": ["start"]
    },
    {
      "name": "Launch React Frontend",
      "type": "node",
      "request": "launch",
      "cwd": "${workspaceFolder}/client",
      "runtimeExecutable": "npm",
      "runtimeArgs": ["start"]
    }
  ],
  "compounds": [
    {
      "name": "Backend + Frontend",
      "configurations": ["Launch Go Backend", "Launch React Frontend"]
    }
  ]
}
```

## Working with the Project in VS Code

### Opening the Project

1. Open VS Code
2. Click on "File" > "Open Folder"
3. Navigate to your NoMercyChain project directory and click "Select Folder"

### Running the Project

#### Running the Backend

1. Open a terminal in VS Code (`Ctrl+` `)
2. Navigate to the project root
3. Run:
   ```powershell
   go run ./cmd/nmxchaind/main.go start
   ```

Or use the VS Code debugger:
1. Press `F5` or click the "Run and Debug" icon in the sidebar
2. Select "Launch Go Backend" from the dropdown

#### Running the Frontend

1. Open a terminal in VS Code
2. Navigate to the client directory:
   ```powershell
   cd client
   ```
3. Run:
   ```powershell
   npm start
   ```

Or use the VS Code debugger:
1. Press `F5` or click the "Run and Debug" icon in the sidebar
2. Select "Launch React Frontend" from the dropdown

### Debugging

#### Debugging the Backend

1. Set breakpoints by clicking in the gutter next to the line numbers
2. Start the debugger with "Launch Go Backend" configuration
3. Use the debug toolbar to step through code, inspect variables, etc.

#### Debugging the Frontend

1. Set breakpoints in your React code
2. Start the debugger with "Launch React Frontend" configuration
3. The Chrome DevTools will open automatically for debugging

## Useful VS Code Shortcuts

- `Ctrl+P`: Quick file navigation
- `Ctrl+Shift+F`: Search across all files
- `F12`: Go to definition
- `Alt+F12`: Peek definition
- `Shift+F12`: Find all references
- `F2`: Rename symbol
- `Ctrl+Space`: Trigger suggestions
- `Ctrl+.`: Quick fixes and refactorings
- `Ctrl+` ``: Toggle terminal
- `Ctrl+B`: Toggle sidebar

## Recommended Workflow

1. Use the Explorer view to navigate the project structure
2. Use the Source Control view to manage Git operations
3. Use the Run and Debug view for debugging
4. Use the Terminal for running commands
5. Use the Problems view to see errors and warnings

## Troubleshooting

### Go Tools Not Working

If Go tools are not working properly:

1. Open the Command Palette (`Ctrl+Shift+P`)
2. Run "Go: Restart Language Server"

### ESLint or Prettier Issues

If ESLint or Prettier are not working:

1. Make sure you have the necessary configuration files in your project:
   - `.eslintrc.js` or `.eslintrc.json` for ESLint
   - `.prettierrc` for Prettier
2. Run `npm install` in the client directory to ensure all dependencies are installed

### Debugging Not Working

If debugging is not working:

1. Make sure you have the latest version of VS Code
2. Make sure you have the latest version of the Go extension
3. Check that the paths in your launch configurations are correct

## Additional Resources

- [VS Code Go Documentation](https://code.visualstudio.com/docs/languages/go)
- [VS Code React Documentation](https://code.visualstudio.com/docs/nodejs/reactjs-tutorial)
- [Debugging in VS Code](https://code.visualstudio.com/docs/editor/debugging)