import type { User } from '$lib/types/user';
import avatar from '$lib/assets/avatar.png';

// temp var
export const user = $state<User>({
	name: 'Admin Verse',
	email: 'admin@sendio.dev',
	isPremium: true,
	avatarUrl: avatar,
	usage: {
		sent: 5863,
		quota: 8000,
		growthPercent: 12.4,
		dailyUsage: [22, 9, 19, 4, 14, 6, 11] // Mon-Sun
	}
});

export function getGrowthDisplay(): string {
	const sign = user.usage.growthPercent >= 0 ? '+' : '';
	return `${sign}${user.usage.growthPercent}%`;
}

export function getUsagePercent(): number {
	return (user.usage.sent / user.usage.quota) * 100;
}

export function getDisplayEmail(): string {
	const [username, domain] = user.email.split('@');
	if (username.length > 15) {
		const first = username.slice(0, 4);
		const last = username.slice(-4);
		return `${first}...${last}@${domain}`;
	}
	return user.email;
}
