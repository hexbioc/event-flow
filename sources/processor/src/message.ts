import type pino from "pino";
import { loggerForModule } from "./logger";

class Message {
	logger: pino.Logger;

	constructor() {
		this.logger = loggerForModule("message");
	}

	get(): string {
		const message = "Hello there!";

		this.logger.info("Sending message <%s>", message);
		return message;
	}
}

export default new Message();
