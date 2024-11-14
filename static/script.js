document.addEventListener('DOMContentLoaded', function(){
    const name = "RAY JONES MUIRURI";
    const writerElement = document.getElementById("name-writer");

    function autoTypeEffect(text, delay=5000) {
        let i = 0;
        writerElement.textContent = "";
        const type = () => {
            if (i < text.length) {
                writerElement.textContent += text[i];
                i++;
                setTimeout(type,100); //type speed
            }else {
                setTimeout(() => {
                    writerElement.textContent = "";
                    i = 0;
                    type();
                }, delay);
            }
        };
        type();
    }

    autoTypeEffect(name);
})