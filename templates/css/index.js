const buttonMonitor = document.getElementById('btn-monitorar')
        const explicacao = document.getElementById('explicacao')
        const resultado = document.getElementById('resultado')
        
        buttonMonitor.addEventListener('click', () =>{
            if(!explicacao.classList.add('explicacao-desativo') & !resultado.classList.add('resultados-ativo')){
                explicacao.classList.remove('explicacao-ativo')
                explicacao.classList.add('explicacao-desativo')
                resultado.classList.remove('resultados-desativo')
                resultado.classList.add('resultados-ativo')
            } else{
                alert("I am an alert box!")
            }
        })