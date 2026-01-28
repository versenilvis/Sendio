<script lang="ts">
	let { class: className = "h-4 w-4" } = $props<{ class?: string }>();
</script>

<svg 
    xmlns="http://www.w3.org/2000/svg" 
    viewBox="0 0 24 24" 
    fill="currentColor" 
    class="{className} animate-pop"
>
    <path d="M8.5 18A2.5 2.5 0 0 1 6 15.5V8H4c-1.1 0-2 .9-2 2v10c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2v-2z"></path>
    <path d="M20 2H10c-1.1 0-2 .9-2 2v10c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2m-1.29 5.71L14 12.42l-2.71-2.71-1-1L11.7 7.3l1 1 1.29 1.29 3.29-3.29 1-1 1.41 1.41-1 1Z"></path>
</svg>

<style>
    @keyframes pop {
        0% {
            opacity: 0;
            transform: scale(0.5) rotate(-10deg);
        }
        50% {
            opacity: 1;
            transform: scale(1.1) rotate(5deg);
        }
        100% {
            opacity: 1;
            transform: scale(1) rotate(0deg);
        }
    }

    .animate-pop {
        animation: pop 0.4s cubic-bezier(0.34, 1.56, 0.64, 1) forwards;
    }
</style>
