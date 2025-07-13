import { config } from "./config";
import type { TargetPayloadDTO } from "./dto";
import { loggerForModule } from "./logger";
import { sleep } from "./utils";

class TargetManager {
	logger = loggerForModule("target");
	bucketSize: number;
	bucketRestoreRate: number;
	invocationCost: number;

	constructor() {
		this.bucketSize = config.TARGET_BUCKET_SIZE;
		this.bucketRestoreRate = config.TARGET_RESTORE_RATE;
		this.invocationCost = config.TARGET_INVOCATION_COST;
	}

	async invoke(payload: TargetPayloadDTO) {
		this.logger.info(`Target invoked with payload: ${JSON.stringify(payload)}`);

		// Emulate processing
		await sleep(Math.random() * 1000);
	}
}

export default new TargetManager();
