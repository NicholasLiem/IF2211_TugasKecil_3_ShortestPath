let visualizer = document.getElementById("visualizer");
let ctx = visualizer.getContext("2d");

class Node {
    constructor(index, name, latitude, longitude) {
        this.index = index;
        this.name = name;
        this.latitude = latitude;
        this.longitude = longitude;
    }
}

class Graph {
    constructor({nodes, edges}) {
        this.nodes = nodes;
        this.edges = edges;
    }
}

function getMousePos(canvas, evt) {
    const rect = canvas.getBoundingClientRect(), // abs. size of element
        scaleX = canvas.width / rect.width,    // relationship bitmap vs. element for x
        scaleY = canvas.height / rect.height;  // relationship bitmap vs. element for y

    return {
        x: (evt.clientX - rect.left) * scaleX,   // scale mouse coordinates after they have
        y: (evt.clientY - rect.top) * scaleY     // been adjusted to be relative to element
    };
}

function drawCircle(x, y, radius) {
    ctx.beginPath();
    ctx.arc(x, y, radius, 0, 2 * Math.PI);
    ctx.stroke();
}

visualizer.style.width = '100%';
visualizer.style.height = '100%';
visualizer.width = visualizer.offsetWidth;
visualizer.height = visualizer.offsetHeight;

// visualizer.addEventListener("mousemove", ev => {
//     const pos = getMousePos(visualizer, ev)
//     ctx.clearRect(0, 0, visualizer.width, visualizer.height)
//     drawCircle(pos.x, pos.y, 20)
// });

window.addEventListener("resize", ev => {
    const width = visualizer.clientWidth;
    const height = visualizer.clientHeight;
    visualizer.width = width;
    visualizer.height = height;
});

function parseFile(file) {
    const reader = new FileReader();
    reader.onload = async function (e) {
        const response = await fetch("http://localhost:8080/parse", {
            method: "POST",
            body: e.target.result
        })
        if (response.status === 204) {
            const error = document.getElementById("error")
            const msg = error.getElementsByClassName("error-msg")[0]
            msg.innerHTML = "Failed to parse file"
            error.style.visibility = "visible"
        } else {
            const error = document.getElementById("error")
            const canvas = document.getElementById("visualizer")
            canvas.style.backgroundColor = "#eee"
            error.style.visibility = "hidden"
            const {Nodes, Edges} = await response.json()
            console.log(new Graph({
                nodes: Nodes,
                edges: Edges,
            }))
        }
    }
    reader.readAsDataURL(file);
    return false;
}
