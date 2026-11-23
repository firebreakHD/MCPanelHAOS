import { Code, ConnectError } from '@connectrpc/connect';
import { rpcClient, silentCallOptions } from './rpc-client';

export async function startOrSwitchServer(id: string, name: string): Promise<boolean> {
	try {
		await rpcClient.server.startServer({ id }, silentCallOptions);
		return true;
	} catch (error) {
		if (!(error instanceof ConnectError) || error.code !== Code.FailedPrecondition) throw error;
		const activeServer = error.metadata.get('X-Mcpanel-Active-Server');
		if (!activeServer) throw error;
		if (!window.confirm(`${error.rawMessage} Welt speichern und zu „${name}“ wechseln?`)) {
			return false;
		}
		await rpcClient.server.startServer({ id, switchFromServerId: activeServer }, silentCallOptions);
		return true;
	}
}
