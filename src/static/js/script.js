const visualizer = document.getElementById("visualizer");
const ctx = visualizer.getContext("2d");
const nodeRadius = 7;
const coords = [];
let graph;
let onhover = -1;

class Node {
    constructor({index, name, latitude, longitude}) {
        this.index = index;
        this.name = name;
        this.latitude = latitude;
        this.longitude = longitude;
    }
}

class Graph {
    constructor({nodes, edges}) {
        this.nodes = [];
        for (let i = 0; i < Object.keys(nodes).length; i++) {
            const item = nodes[i];
            this.nodes.push(new Node({
                index: item["Index"],
                name: item["Name"],
                latitude: item["Latitude"],
                longitude: item["Longitude"],
            }));
        }
        this.maxLat = this.nodes.reduce((prev, curr) => prev.latitude > curr.latitude ? prev : curr).latitude;
        this.minLat = this.nodes.reduce((prev, curr) => prev.latitude < curr.latitude ? prev : curr).latitude;
        this.maxLon = this.nodes.reduce((prev, curr) => prev.longitude > curr.longitude ? prev : curr).longitude;
        this.minLon = this.nodes.reduce((prev, curr) => prev.longitude < curr.longitude ? prev : curr).longitude;
        this.edges = edges;
    }

    draw(canvas) {
        const ctx = canvas.getContext("2d");
        const width = canvas.width * 0.7;
        const height = canvas.height * 0.7;
        for (const [idx, node] of this.nodes.entries()) {
            const x = canvas.width * 0.15 + (node.longitude - this.minLon) / (this.maxLon - this.minLon) * width;
            const y = canvas.height * 0.15 + (node.latitude - this.minLat) / (this.maxLat - this.minLat) * height;
            coords[idx] = [x, y];
        }
        const drawn = [];
        for (const [from, edges] of Object.entries(this.edges)) {
            for (const [to, weight] of Object.entries(edges)) {
                if (!([from, to] in drawn)) {
                    const [x1, y1] = coords[from];
                    const [x2, y2] = coords[to];
                    drawLine(x1, y1, x2, y2, ctx);
                    drawText(weight.toString(), (x1 + x2) / 2 - 10, (y1 + y2) / 2 - 10, ctx);
                    drawn.push([from, to]);
                    drawn.push([to, from]);
                }
            }
        }
        for (const [idx, [x, y]] of coords.entries()) {
            const radius = nodeRadius + (onhover === idx ? 3 : 0);
            drawText(this.nodes[idx].name, x, y - radius - 10, ctx);
            drawCircle(x, y, radius, ctx);
        }
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

function drawCircle(x, y, radius, ctx) {
    ctx.beginPath();
    ctx.fillStyle = "#fff"
    ctx.arc(x, y, radius, 0, 2 * Math.PI);
    ctx.fill();
}

function drawText(text, x, y, ctx) {
    ctx.font = "15px Poppins";
    ctx.textAlign = "center";
    ctx.fillStyle = "#fff"
    ctx.fillText(text, x, y)
}

function drawLine(x1, y1, x2, y2, ctx) {
    ctx.beginPath();
    ctx.strokeStyle = "#a0aaef"
    ctx.moveTo(x1, y1);
    ctx.lineTo(x2, y2);
    ctx.stroke();
}

visualizer.style.width = '100%';
visualizer.style.height = '100%';
visualizer.width = visualizer.offsetWidth;
visualizer.height = visualizer.offsetHeight;

visualizer.addEventListener("mousemove", ev => {
    const pos = getMousePos(visualizer, ev)
    if (graph) {
        for (const [idx, [x, y]] of coords.entries()) {
            const radius = nodeRadius + (onhover === idx ? 5 : 0);
            if (pos.x >= x - radius && pos.x <= x + radius && pos.y >= y - radius && pos.y <= y + radius) {
                visualizer.style.cursor = "pointer";
                if (onhover !== idx) {
                    onhover = idx;
                    ctx.clearRect(0, 0, visualizer.width, visualizer.height);
                    graph.draw(visualizer);
                }
                break;
            } else {
                if (onhover !== -1) {
                    onhover = -1;
                    ctx.clearRect(0, 0, visualizer.width, visualizer.height);
                    graph.draw(visualizer);
                }
                visualizer.style.cursor = "default";
            }
        }
    }
});

window.addEventListener("resize", _ => {
    const width = visualizer.clientWidth;
    const height = visualizer.clientHeight;
    visualizer.width = width;
    visualizer.height = height;
    graph.draw(visualizer)
});

function displayError(message) {
    const error = document.getElementById("error")
    const msg = error.getElementsByClassName("error-msg")[0]
    msg.innerHTML = message
    error.style.visibility = "visible"
}

function hideError() {
    const error = document.getElementById("error")
    const canvas = document.getElementById("visualizer")
    canvas.style.backgroundColor = "#111"
    error.style.visibility = "hidden"
}

function parseFile(file) {
    const reader = new FileReader();
    reader.onload = (e) => {
        fetch("http://localhost:8080/parse", {
            method: "POST",
            body: e.target.result
        }).then((response) => {
            if (response.ok) {
                if (response.status === 204) {
                    throw new Error("Failed to parse file")
                }
                return response.json()
            }
            throw new Error("Failed to upload file")
        }).then(async (responseJson) => {
            hideError()
            const {Nodes, Edges} = await responseJson
            graph = new Graph({
                nodes: Nodes,
                edges: Edges,
            })
            console.log(graph)
            ctx.clearRect(0, 0, visualizer.width, visualizer.height)
            graph.draw(visualizer)
        }).catch((error) => {
            displayError(error.message)
        })
    }
    reader.readAsDataURL(file);
    return false;
}

window.onload = function () {
    Particles.init({
        selector: ".background",
        color: "#a0aaef",
        connectParticles: true,
    });
};