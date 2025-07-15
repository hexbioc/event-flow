export interface EventDTO {
	id: string;
	name: string;
	body: string;
	timestamp: string;
}

export interface TargetPayloadDTO extends EventDTO {
	brand: string;
}
