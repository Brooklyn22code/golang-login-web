let img_place = document.getElementById("img") 

let next_btn = document.getElementById("next")
let up_btn = document.getElementById("up")
let font_change = document.getElementById("change_font")

let page_num = document.getElementById("page_num")

let text = document.getElementById("background")


const imgs = ["/statics/imgs/HTML_img.jpg", "/statics/imgs/golang-logo.webp", "/statics/imgs/css_img.png", "/statics/imgs/python_img.jpg"];
const fonts = ["Fantasy", "Monospace", "Cursive", "Serif", 'Sans-serif']

const texts = ["Hypertext Markup Language (HTML) is the standard markup language for documents designed to be displayed in a web browser. It defines the content and structure of web content. It is often assisted by technologies such as Cascading Style Sheets (CSS) and scripting languages such as JavaScript.",
    "Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency. It is often referred to as Golang because of its former domain name.",
    "Cascading Style Sheets (CSS) is a style sheet language used for specifying the presentation and styling of a document written in a markup language such as HTML or XML (including XML dialects such as SVG, MathML or XHTML). CSS is a cornerstone technology of the World Wide Web, alongside HTML and JavaScript.",
    "Python is a high-level, general-purpose programming language. You can use python in very more purposes, like make a AI, make a website backend by using Django or flask, data Ana."

]

function color(x) {
    document.body.style.backgroundColor = x
}

function random() {
    let red = Math.floor(Math.random() * 255)
    let blue = Math.floor(Math.random() * 255)
    let yellow = Math.floor(Math.random() * 255)

    document.body.style.backgroundColor = `rgb(${red} ${yellow} ${blue})`
}

var index_list = 0



up_btn.addEventListener('click', () => {
    if (index_list > 0) {
        index_list -= 1
    } else {
        index_list = imgs.length - 1
    }
    img_place.src = imgs[index_list]
    page_num.innerHTML = `${index_list + 1} / ${imgs.length}`
    text.innerHTML = texts[index_list] 
})


next_btn.addEventListener('click', () => {
    if (imgs.length - 1 > index_list) {
        index_list += 1
    } else {
        index_list = 0
    }
    img_place.src = imgs[index_list]
    page_num.innerHTML = `${index_list + 1} / ${imgs.length}`
    text.innerHTML = texts[index_list] 
})

var index_fonts = 0

font_change.addEventListener('click', () => {
    if (fonts.length - 1 > index_fonts) {
        index_fonts += 1
    } else {
        index_fonts = 0
   }
    text.style.fontFamily = fonts[index_fonts]
})
