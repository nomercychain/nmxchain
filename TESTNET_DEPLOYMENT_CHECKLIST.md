# NoMercyChain Testnet Deployment Checklist

This checklist ensures that all components are properly configured and ready for testnet deployment.

## Backend Preparation

- [x] Chain binary builds successfully
- [x] All modules are properly registered in the app
- [x] Genesis file is properly configured with all modules
- [x] Testnet setup script works correctly
- [x] Chain can be initialized and started
- [x] REST API endpoints are accessible
- [x] CORS is enabled for frontend access
- [x] Swagger documentation is enabled

## Frontend Preparation

- [x] All required dependencies are installed
- [x] API client is configured with correct endpoints
- [x] Transaction message types are registered correctly
- [x] Environment variables are set up for testnet
- [x] Frontend can connect to the backend
- [x] Keplr wallet integration is configured

## Deployment Scripts

- [x] Build script works correctly
- [x] Testnet setup script works correctly
- [x] Reset testnet script works correctly
- [x] Monitoring script works correctly
- [x] Deployment package script works correctly

## Testing

- [x] Chain starts successfully
- [x] Accounts can be created
- [x] Tokens can be sent between accounts
- [x] DynaContract module functions correctly
- [x] DeAI module functions correctly
- [x] HyperChain module functions correctly
- [x] TruthGPT module functions correctly
- [x] Frontend connects to the backend
- [x] Transactions can be signed and broadcast

## Deployment Steps

1. **Build the Chain**
   ```
   ./scripts/build.ps1
   ```

2. **Deploy the Testnet**
   ```
   ./scripts/deploy_testnet.ps1
   ```

3. **Start the Full Testnet (Backend and Frontend)**
   ```
   ./scripts/start_testnet_full.ps1
   ```

4. **Monitor the Testnet**
   ```
   ./scripts/monitor_testnet.ps1 -Continuous
   ```

5. **Test the Modules**
   ```
   ./scripts/test_dynacontract.ps1
   ./scripts/test_integration.ps1
   ```

## Troubleshooting

### Common Issues and Solutions

1. **Chain fails to start**
   - Check logs for errors
   - Ensure ports 26656, 26657, and 1317 are available
   - Try resetting the testnet with `./scripts/reset_testnet.ps1`

2. **Frontend can't connect to backend**
   - Verify the chain is running
   - Check CORS settings in the chain's config.toml
   - Verify API endpoints in the frontend configuration

3. **Transaction errors**
   - Verify account has sufficient balance
   - Check that the chain-id is correct
   - Ensure the keyring backend is set to "test"

4. **Module errors**
   - Check module registration in app.go
   - Verify module parameters in genesis.json
   - Check module handler implementation

## Post-Deployment Verification

- [ ] Chain is running and producing blocks
- [ ] REST API is accessible
- [ ] Frontend is accessible
- [ ] Accounts can be created and funded
- [ ] All modules are functioning correctly
- [ ] Transactions can be signed and broadcast
- [ ] Monitoring is working correctly