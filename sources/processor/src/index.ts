import { config } from "./config";
import { createHealthcheckServer } from "./health";
import { loggerForModule } from "./logger";
import rmqClient from "./rmq-client";
import { QueueWorker } from "./worker";
import type net from "node:net";

async function main() {
	const logger = loggerForModule("main");
	let workers: QueueWorker[] = [];
	let healthServer: net.Server | undefined;

	async function shutdownGracefully() {
		// Stop workers and close channels
		try {
			await Promise.all(
				workers.map((worker) => worker.stop().catch((_) => {})),
			);
		} catch {}

		// Close RabbitMQ connection
		try {
			await rmqClient.close();
		} catch {}

		// Stop healthcheck server
		if (healthServer) healthServer.close();
	}

	logger.info(`Starting application process [${process.pid}]`);
	try {
		// Connect to RabbitMQ
		await rmqClient.connect();
		logger.info("Connection established to RabbitMQ");

		// Ensure RabbitMQ is setup correctly
		await rmqClient.assertQueue(config.RMQ_QUEUE, { durable: true });
		logger.info("RabbitMQ checks completed");

		// Setup and start workers
		workers = await Promise.all(
			Array.from(Array(config.WORKERS)).map(async (_, index) => {
				const workerName = `worker-${String(index).padStart(3, "0")}`;
				const workerChannel = await rmqClient.newChannel();

				const worker = new QueueWorker(
					workerName,
					config.RMQ_QUEUE,
					workerChannel,
				);
				worker.start();

				return worker;
			}),
		);

		// Start healthcheck server
		healthServer = createHealthcheckServer();
		healthServer.listen(config.HEALTHCHECK_PORT);
	} catch (err) {
		logger.error(err, "Application failed to boot!");
		shutdownGracefully();
	}
	logger.info("Workers created and started successfully");

	// Attach SIGINT and SIGTERM listeners
	process.on("SIGINT", () => {
		logger.info("SIGINT received, closing all connections and shutting down");
		shutdownGracefully();
	});
	process.on("SIGTERM", () => {
		logger.info("SIGTERM received, closing all connections and shutting down");
		shutdownGracefully();
	});
}

main();
