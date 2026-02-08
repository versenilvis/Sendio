<script lang="ts">
</script>

<div class="bg-black rounded-[44px] p-7 w-48 h-48 shadow-2xl flex flex-col justify-between group transition-all duration-700 select-none overflow-hidden relative">
    
    <!-- Massive Liquid Lightning Bolt -->
    <div class="absolute right-[-45%] bottom-[-12%] w-[120%] h-[120%] pointer-events-none">
        <svg viewBox="0 0 24 24" class="w-full h-full">
            <defs>
                <clipPath id="boltClipSharp">
                    <path d="M19 9h-5V3a1 1 0 0 0-.69-.95c-.41-.13-.86.01-1.12.36l-8 11a1 1 0 0 0-.08 1.04A1 1 0 0 0 5 15h5v6a1 1 0 0 0 1 1c.31 0 .62-.15.81-.41l8-11a1 1 0 0 0 .08-1.04A1 1 0 0 0 19 9z" />
                </clipPath>
            </defs>
            
            <g clip-path="url(#boltClipSharp)">
                <!-- Lighter background for the empty part of bolt -->
                <rect x="0" y="0" width="24" height="24" fill="#222" />
                
                <!-- Yellow liquid with a "shorter" / more compact wave -->
                <g style="filter: drop-shadow(0 -1px 1px rgba(254, 240, 7, 0.8))">
                    <path 
                        fill="#FEF007" 
                        d="M0,11.5 C4,10 6,13 12,12 C18,11 22,12 22,11.5 V24 H0 Z"
                    />
                </g>
            </g>
        </svg>
    </div>

    <!-- Content -->
    <div class="relative z-10 flex flex-col pt-1">
        <div class="flex items-baseline gap-1">
            <span class="text-4xl font-bold text-[#FEF007] tracking-tighter drop-shadow-[0_0_10px_rgba(254,240,7,0.8)]">54.2<span class="text-2xl ml-0.5">%</span></span>
        </div>
        <span class="text-[11px] font-bold text-white/60 tracking-tight mt-1 uppercase">Newsletter</span>
    </div>

    <div class="relative z-10 flex flex-col mb-1">
        <span class="text-xl font-bold text-white tracking-tight leading-tight">Subscribers</span>
    </div>
</div>
