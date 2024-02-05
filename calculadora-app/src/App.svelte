<script>
	
	import BotonesNum from "./BotonesNum.svelte";
    import Operaciones from "./Operaciones.svelte";
    import Historial from "./Historial.svelte";
    import './styles.css';
	
	let inputValue = '';
	let error ='';

	const handleButtonClick =  async (value) => {
		if (value === '=') {
      handleEquals();
    } else if (value === 'C' || value === 'CE') {
      inputValue = '';
      error = '';
    } else {
      inputValue += value;
    }
   /*ERROR DE ESCRITURA DEL IF DENTRO DEL ASYNC NO SE PORQUE
   {#if value === '='} 
	handleButtonClick();		
    {:else if value === 'C'}
      inputValue = ''; 
	  error = '';
    {:else if value === 'CE'}
      inputValue = '';
	  error = '';
    {:else}
      inputValue += value; 
    {/if}*/
  };

  const handleEquals = async () => {
	try {
      const response = await fetch(`http://localhost:8000/calculadora?valor=${inputValue}`);
      const data = await response.json();
      
      if (data.resultado !== undefined) {
        if (data.error) {
          error = `Error al realizar la operación: ${data.error}`;
        } else {
          if (typeof data.resultado === 'number' && !isNaN(data.resultado) && isFinite(data.resultado)) {
            inputValue = (data.resultado % 1 !== 0) ? data.resultado.toFixed(2) : data.resultado;
          }
        }
      }
    } catch (error) {
      console.error('Error al procesar la respuesta:', error.message);
      error = `Error al procesar la respuesta: ${error.message}`;
    }

  };

</script>

<!-- componente-->
<div>
	<input bind:value={inputValue} class="p-2 mb-4" />
	<BotonesNum on:click={handleButtonClick} />
	<div class="p-2 mt-4 bg-black text-white">
		{#if error}
		{error}
		{:else}
		 Resultado: {inputValue}
		 {/if}

	</div>
  </div>

