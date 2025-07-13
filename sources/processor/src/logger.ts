import pino from "pino";

const _baseLogger = pino();

export function loggerForModule(module: string) {
	return _baseLogger.child({ module });
}
