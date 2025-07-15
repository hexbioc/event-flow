import type amqp from "amqplib";
import { sleep } from "./utils";
import { loggerForModule } from "./logger";
import type pino from "pino";
import type { EventDTO } from "./dto";
import target from "./target";

const pollIntervalMillis = 1000;

export class QueueWorker {
	name: string;
	queue: string;
	channel: amqp.Channel;
	logger: pino.Logger;
	isActive = false;

	constructor(name: string, queue: string, channel: amqp.Channel) {
		this.name = name;
		this.queue = queue;
		this.channel = channel;
		this.logger = loggerForModule(this.name);
	}

	async start() {
		this.logger.info(`Starting worker <${this.name}>`);

		this.isActive = true;
		while (this.isActive) {
			// Fetch a message from the queue
			const message = await this.channel.get(this.queue);

			if (message) {
				// Process message
				await this.handleMessage(message);
			} else {
				// Wait before checking again
				await sleep(pollIntervalMillis);
			}
		}
	}

	async handleMessage(message: amqp.GetMessage) {
		const event: EventDTO = JSON.parse(message.content.toString());

		// Modify event as required
		const targetPayload = { ...event, brand: "testBrand" };

		// Invoke target
		try {
			await target.invoke(targetPayload);
		} catch (err) {
			this.logger.error(err, "Failure in processing event", {
				payload: JSON.stringify(targetPayload),
			});
			await this.channel.nack(message);
		}

		await this.channel.ack(message);
	}

	async stop() {
		this.isActive = false;
		await this.channel.close();
	}
}
