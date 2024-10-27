package main

// layout/fa-base-layout.html

const fa_html_start = `
<!DOCTYPE html>
<html>
`
const fa_head_start = `
<head>
    <meta charset='utf-8'>
    <meta http-equiv='X-UA-Compatible' content='IE=edge'>
    <meta name='viewport' content='width=device-width, initial-scale=1'>

    <!-- PreLoads -->
    <link rel="preload" href="/public/fonts/Vazirmatn[wght].woff2" as="font" type="font/woff2" crossorigin />


    <!-- Fonts -->

    <!-- Stylesheets -->
    <link rel="stylesheet" type='text/css' media='screen' href='/public/styles/main.css' />

    <!-- Scripts -->
    <script src='/public/libs/htmx.min.js'></script>
    <script src='/public/libs/htmx.preload.js'></script>
`

// <link rel="preload" href='/public/styles/main.css' as="style" type='text/css' media='screen' crossorigin />
// <link rel="preload" href="/public/libs/htmx.min.js" as="script" crossorigin />
// <link rel="preload" href="/public/libs/htmx.preload.js" as="script" crossorigin />

const fa_head_end = `
</head>
`
const fa_body_start = `
<body id="body-box" class="fa" hx-ext="preload">
`
const fa_shell_start = `
    <div id="shell">
`
const fa_shell_end = `
    </div>
`
const fa_turtle_start = `
    <div id="turtle">
`
const fa_turtle_end = `
    </div>
`
const fa_body_end = `
</body>
`
const fa_html_end = `
</html>
`

const en_html_start = `
<!DOCTYPE html>
<html>
`
const en_head_start = `
<head>
    <meta charset='utf-8'>
    <meta http-equiv='X-UA-Compatible' content='IE=edge'>
    <meta name='viewport' content='width=device-width, initial-scale=1'>

    <!-- PreLoads -->
    <link rel="preload" href="/public/fonts/Vazirmatn[wght].woff2" as="font" type="font/woff2" crossorigin />


    <!-- Fonts -->

    <!-- Stylesheets -->
    <link rel="stylesheet" type='text/css' media='screen' href='/public/styles/main.css' />

    <!-- Scripts -->
    <script src='/public/libs/htmx.min.js'></script>
    <script src='/public/libs/htmx.preload.js'></script>
`

// <link rel="preload" href='/public/styles/main.css' as="style" type='text/css' media='screen' crossorigin />
// <link rel="preload" href="/public/libs/htmx.min.js" as="script" crossorigin />
// <link rel="preload" href="/public/libs/htmx.preload.js" as="script" crossorigin />

const en_head_end = `
</head>
`
const en_body_start = `
<body id="body-box" class="en" hx-ext="preload">
`
const en_shell_start = `
    <div id="shell">
`
const en_shell_end = `
    </div>
`
const en_turtle_start = `
    <div id="turtle">
`
const en_turtle_end = `
    </div>
`
const en_body_end = `
</body>
`
const en_html_end = `
</html>
`

const fa_base_shell = `
<div id="shell-header">
    <p><a href="/">تربتی</a></p>
</div>
`
const en_base_shell = `
<div id="shell-header">
    <p><a href="/en">Torbatti</a></p>
</div>
`

// <p hx-get="/hx/fa/content/zbnc" hx-trigger="click" hx-target="#turtle" hx-swap="innerHtml" preload="preload:init"> زبان برنامه نویسی سی</p>
