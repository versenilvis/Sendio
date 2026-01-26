<script lang="ts">
	import { onMount } from 'svelte';
	import { fly, fade } from 'svelte/transition';
	import MenuIcon from './MenuIcon.svelte';
	import MobileMenu from './MobileMenu.svelte';

	let open = $state(false);
	let scrolled = $state(false);
	let activeMenu = $state<string | null>(null);

	// handle scroll state
	onMount(() => {
		const handleScroll = () => {
			scrolled = window.scrollY > 10;
		};
		window.addEventListener('scroll', handleScroll);
		handleScroll();
		return () => window.removeEventListener('scroll', handleScroll);
	});

	// stop body scroll when mobile menu is open
	$effect(() => {
		if (open) {
			document.body.style.overflow = 'hidden';
		} else {
			document.body.style.overflow = '';
		}
	});

	const products = [
		{ title: 'Website Builder', desc: 'Create responsive sites with ease', icon: 'bx-globe' },
		{ title: 'Cloud Platform', desc: 'Deploy and scale apps in cloud', icon: 'bx-layer' },
		{ title: 'Collaboration', desc: 'Tools for better teamwork', icon: 'bx-user-plus' },
		{ title: 'Analytics', desc: 'Track and analyze traffic', icon: 'bx-bar-chart-alt-2' },
		{ title: 'Integrations', desc: 'Connect your apps and services', icon: 'bx-plug' },
		{ title: 'API', desc: 'Build custom integrations', icon: 'bx-code-block' }
	];

	const company = [
		{ title: 'About Us', desc: 'Learn more about our story', icon: 'bx-group' },
		{ title: 'Customer Stories', desc: 'How we help clients succeed', icon: 'bx-star' },
		{ title: 'Partnerships', desc: 'Collaborate for mutual growth', icon: 'bx-handicap' } // using generic group icon for partnerships
	];

	const legalAndOther = [
		{ title: 'Terms of Service', icon: 'bx-file' },
		{ title: 'Privacy Policy', icon: 'bx-shield-quarter' },
		{ title: 'Refund Policy', icon: 'bx-refresh' },
		{ title: 'Blog', icon: 'bx-leaf' },
		{ title: 'Help Center', icon: 'bx-help-circle' }
	];

	function toggleMenu(menu: string) {
		if (activeMenu === menu) activeMenu = null;
		else activeMenu = menu;
	}

	function closeMenus() {
		activeMenu = null;
	}
</script>

<!-- navigation bar -->
<header
	class="fixed top-0 z-50 w-full transition-all duration-300 {scrolled ? 'glass border-b border-(--color-border)/50 h-14' : 'h-16'}"
>
	<nav class="mx-auto flex h-full max-w-5xl items-center justify-between px-6">
		<div class="flex items-center gap-10">
			<!-- logo -->
			<a href="/" class="flex items-center gap-2.5 text-xl font-bold tracking-tight text-black">
				<div class="flex h-8 w-8 items-center justify-center rounded-lg bg-blue-600/5 text-blue-600">
					<i class="bx bx-paper-plane"></i>
				</div>
				<span>Sendio</span>
				<span class="rounded-md bg-blue-50 px-1.5 py-0.5 text-[9px] font-black uppercase text-blue-600 tracking-wider">Experimental</span>
			</a>

			<!-- desktop menu -->
			<div class="hidden items-center gap-2 text-[14px] font-semibold md:flex">
				<!-- product dropdown -->
				<div 
					class="relative h-16 flex items-center"
					onmouseenter={() => activeMenu = 'product'}
					onmouseleave={closeMenus}
				>
					<button 
						class="flex items-center gap-1.5 rounded-xl px-4 py-2 transition hover:bg-(--color-bg-secondary)/50 {activeMenu === 'product' ? 'bg-(--color-bg-secondary)/50' : ''}"
					>
						Product <i class="bx bx-chevron-down text-lg transition {activeMenu === 'product' ? 'rotate-180' : ''}"></i>
					</button>

					{#if activeMenu === 'product'}
						<div 
							transition:fly={{ y: 10, duration: 200 }}
							class="absolute top-[calc(100%-8px)] left-0 w-[580px]"
						>
							<div class="iphone-card overflow-hidden border border-(--color-border)/50 bg-white shadow-2xl">
								<div class="grid grid-cols-2 gap-2 p-3">
									{#each products as item}
										<a href="#" class="flex items-center gap-4 rounded-2xl p-3 transition hover:bg-(--color-bg-secondary)/50">
											<div class="flex h-12 w-12 items-center justify-center rounded-xl bg-(--color-bg-secondary) text-2xl text-(--color-text-primary)">
												<i class="bx {item.icon}"></i>
											</div>
											<div class="flex flex-col">
												<span class="font-bold text-(--color-text-primary)">{item.title}</span>
												<span class="text-xs text-(--color-text-muted)">{item.desc}</span>
											</div>
										</a>
									{/each}
								</div>
								<!-- demo block -->
								<div class="bg-(--color-bg-secondary)/30 border-t border-(--color-border)/30 p-4">
									<p class="text-[13px] text-(--color-text-muted)">
										Interested? <a href="#" class="font-bold text-(--color-text-primary) hover:underline">Schedule a demo</a>
									</p>
								</div>
							</div>
						</div>
					{/if}
				</div>

				<!-- company dropdown -->
				<div 
					class="relative h-16 flex items-center"
					onmouseenter={() => activeMenu = 'company'}
					onmouseleave={closeMenus}
				>
					<button 
						class="flex items-center gap-1.5 rounded-xl px-4 py-2 transition hover:bg-(--color-bg-secondary)/50 {activeMenu === 'company' ? 'bg-(--color-bg-secondary)/50' : ''}"
					>
						Company <i class="bx bx-chevron-down text-lg transition {activeMenu === 'company' ? 'rotate-180' : ''}"></i>
					</button>

					{#if activeMenu === 'company'}
						<div 
							transition:fly={{ y: 10, duration: 200 }}
							class="absolute top-[calc(100%-8px)] left-0 w-[500px]"
						>
							<div class="iphone-card overflow-hidden border border-(--color-border)/50 bg-white shadow-2xl">
								<div class="flex">
									<!-- left side -->
									<div class="w-1/2 border-r border-(--color-border)/30 p-3 space-y-1">
										{#each company as item}
											<a href="#" class="flex items-center gap-4 rounded-2xl p-3 transition hover:bg-(--color-bg-secondary)/50">
												<div class="flex h-12 w-12 items-center justify-center rounded-xl bg-(--color-bg-secondary) text-2xl text-(--color-text-primary)">
													<i class="bx {item.icon}"></i>
												</div>
												<div class="flex flex-col">
													<span class="font-bold text-(--color-text-primary)">{item.title}</span>
													<span class="text-xs text-(--color-text-muted)">{item.desc}</span>
												</div>
											</a>
										{/each}
									</div>
									<!-- right side -->
									<div class="w-1/2 p-3 space-y-1">
										{#each legalAndOther as item}
											<a href="#" class="flex items-center gap-3 rounded-xl p-2.5 transition hover:bg-(--color-bg-secondary)/50">
												<i class="bx {item.icon} text-lg text-(--color-text-muted)"></i>
												<span class="font-bold text-(--color-text-primary)">{item.title}</span>
											</a>
										{/each}
									</div>
								</div>
							</div>
						</div>
					{/if}
				</div>

				<a href="#pricing" class="px-4 py-2 transition hover:text-(--color-accent)">Pricing</a>
			</div>
		</div>

		<!-- actions -->
		<div class="flex items-center gap-3">
			<div class="hidden items-center gap-3 md:flex lg:gap-4">
				<button class="text-[14px] font-bold text-(--color-text-primary) transition hover:opacity-70">Sign In</button>
				<a href="/dashboard" class="rounded-full bg-(--color-text-primary) px-5 py-2 text-[14px] font-bold text-white transition hover:opacity-90 active:scale-95 shadow-sm">
					Get Started
				</a>
			</div>

			<!-- mobile toggle -->
			<button 
				class="flex h-10 w-10 items-center justify-center rounded-2xl border border-(--color-border) bg-(--color-bg-secondary)/20 md:hidden"
				onclick={() => open = !open}
			>
				<MenuIcon {open} class="h-6 w-6" />
			</button>
		</div>
	</nav>
</header>

<!-- mobile menu overlay -->
<MobileMenu {open}>
	<div class="flex flex-col gap-6">
		<div class="space-y-4">
			<span class="text-[12px] font-bold uppercase tracking-wider text-(--color-text-muted)">Product</span>
			<div class="grid grid-cols-1 gap-1">
				{#each products as item}
					<a href="#" class="flex items-center gap-4 py-2">
						<div class="flex h-10 w-10 items-center justify-center rounded-xl bg-(--color-bg-secondary) text-xl text-(--color-text-primary)">
							<i class="bx {item.icon}"></i>
						</div>
						<span class="font-bold text-lg">{item.title}</span>
					</a>
				{/each}
			</div>
		</div>
		
		<div class="space-y-4">
			<span class="text-[12px] font-bold uppercase tracking-wider text-(--color-text-muted)">Company</span>
			<div class="grid grid-cols-1 gap-1">
				{#each [...company, ...legalAndOther] as item}
					<a href="#" class="flex items-center gap-4 py-1">
						<i class="bx {item.icon} text-2xl text-(--color-text-muted)"></i>
						<span class="font-bold text-lg">{item.title}</span>
					</a>
				{/each}
			</div>
		</div>

		<a href="#pricing" class="text-xl font-bold border-t border-(--color-border)/20 pt-6">Pricing</a>
	</div>
</MobileMenu>
