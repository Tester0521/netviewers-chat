const numDivs = 20; 
const divs = [];
const block = document.body.querySelector(".circles")

for (let i = 0; i < numDivs; i++) {
    const div = document.createElement('div');

    div.className = (Math.random() * 3 >= 1) ? 'c2' : (Math.random() * 3 >= 2) ? 'c1' : 'c3';


    block.appendChild(div);

    divs.push({
        element: div,
        x: Math.random() * (window.innerWidth - 50),
        y: Math.random() * (window.innerHeight - 50),
        dx: (Math.random() - 1), 
        dy: (Math.random() - 1)  
    });
}

function animate() {
    divs.forEach(div => {
        div.x += div.dx;
        div.y += div.dy;

        if (div.x <= 0 || div.x >= window.innerWidth - 50) div.dx *= -1; 
        if (div.y <= 0 || div.y >= window.innerHeight - 50) div.dy *= -1;

        div.element.style.transform = `translate(${div.x}px, ${div.y}px)`;
    });

    setTimeout(() => requestAnimationFrame(animate), 50)
}

animate(); 

window.addEventListener('resize', () => {
    divs.forEach(div => {
        div.x = Math.min(div.x, window.innerWidth - 50);
        div.y = Math.min(div.y, window.innerHeight - 50);
    });
});