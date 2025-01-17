<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Ramada Front</title>
    @include('components.head-scripts')
    @yield('head')
</head>
<body>
    <main>
        <header>
            @include('.components.menu')
            @yield('header')
        </header>

        <section>
            @yield('body')
        </section>

        <footer>
            <figure>
                <img src="{{ asset('/img/logo.png') }}" alt="logo">
            </figure>

            @yield('footer')
        </footer>
    </main>
</body>
</html>