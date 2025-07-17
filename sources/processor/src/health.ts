import net from "node:net";

export function createHealthcheckServer() {
	return net.createServer((socket) => {
		socket.on("data", (data) => {
			const msg = data.toString().trim();

			if (msg === "ping") {
				socket.write("pong");
			}
		});
	});
}
