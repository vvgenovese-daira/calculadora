<script>
    export let onButtonClick; //manejo de clicks
   
    const handleButtonClick = (value) => {
        onButtonClick(value);
    }
</script>

<!-- Componentes-->
<!-- primer intento
<div class="grid grid-cols-3 gap-2">
    {#each [7, 8, 9, 4, 5, 6, 1, 2, 3, 0, '.', '=', 'C', 'CE'] as buttonValue}
     <button on:click={() => handleButtonClick(buttonValue)} class="bg-blue-500 text-white p-2 rounded">
      {buttonValue}
     </button>
    {/each}
</div>-->

<!-- Componentes-->
<div class="grid grid-cols-4 gap-2">
    {#each ['C', 'CE', '7', '8', '9', '+', '4', '5', '6', '-', '1', '2', '3', '*', '0', '.', '/', '=', '+-'] as buttonValue}
        {#if ['C', 'CE'].includes(buttonValue)}
            <button on:Buttonclick={() => handleButtonClick(buttonValue)} class="bg-orange-500 text-white p-2 rounded">
                {buttonValue}
            </button>
        {:else if buttonValue === '='}
            <button on:Buttonclick={() => handleButtonClick(buttonValue)} class="bg-green-500 text-white p-2 rounded">
                {buttonValue}
            </button>
        {:else if ['+', '-', '*', '/', '%', '+-'].includes(buttonValue)}
            <button on:Buttonclick={() => handleButtonClick(buttonValue)} class="bg-red-700 text-white p-2 rounded">
                {buttonValue}
            </button>
        {:else}
            <button on:Buttonclick={() => handleButtonClick(buttonValue)} class="bg-blue-500 text-white p-2 rounded">
                {buttonValue}
            </button>
        {/if}
    {/each}
</div>

