<script>
    export let onButtonClick; //manejo de clicks
   
    const handleButtonClick = (value) => {
        onButtonClick(value);
    }
</script>

<!-- Componentes-->
<div class="grid grid-cols-3 gap-2">
    {#each [7, 8, 9, 4, 5, 6, 1, 2, 3, 0, '.', '=', 'C', 'CE'] as buttonValue}
     <button on:click={() => handleButtonClick(buttonValue)} class="bg-blue-500 text-white p-2 rounded">
      {buttonValue}
     </button>
    {/each}
</div>
