export interface UserUsage {
	sent: number;
	quota: number;
	growthPercent: number;
	dailyUsage: number[]; // last 7 days
}

export interface User {
	name: string;
	email: string;
	isPremium: boolean;
	avatarUrl: string;
	usage: UserUsage;
}
