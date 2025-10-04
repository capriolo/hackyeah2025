<script>

  import { link, navigate } from "svelte5-router";
</script>

<nav>
    <div class="container">
        <ul>
            <li><a use:link href="/" aria-label="home"><svg width="100%" height="100%" viewBox="0 0 202 202" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" xml:space="preserve" xmlns:serif="http://www.serif.com/" style="fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:2;"><path d="M150.938,201.251l-100.625,-0c-27.769,-0 -50.313,-22.544 -50.313,-50.313l-0,-100.625c-0,-27.769 22.544,-50.313 50.313,-50.313l100.625,0c27.769,0 50.313,22.544 50.313,50.313l-0,100.625c-0,27.769 -22.544,50.313 -50.313,50.313Z" style="fill:url(#_Linear1);"/><path d="M132.579,26.616c3.312,2.393 4.536,6.736 3.033,10.521l-21.821,54.582l40.275,-0c3.758,-0 7.098,2.338 8.378,5.873c1.28,3.534 0.195,7.487 -2.672,9.881l-80.161,66.8c-3.145,2.617 -7.626,2.756 -10.939,0.362c-3.312,-2.393 -4.536,-6.736 -3.033,-10.521l21.821,-54.582l-40.275,0c-3.758,0 -7.098,-2.338 -8.378,-5.873c-1.28,-3.535 -0.195,-7.487 2.672,-9.881l80.161,-66.8c3.145,-2.617 7.626,-2.756 10.939,-0.362Z" style="fill:#fff;fill-rule:nonzero;"/><defs><linearGradient id="_Linear1" x1="0" y1="0" x2="1" y2="0" gradientUnits="userSpaceOnUse" gradientTransform="matrix(1.23231e-14,201.251,-201.251,1.23231e-14,100.625,0)"><stop offset="0" style="stop-color:#ef3d23;stop-opacity:1"/><stop offset="1" style="stop-color:#fbd3d4;stop-opacity:1"/></linearGradient></defs></svg></a></li>
            <li><a use:link href="/profile"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><!--!Font Awesome Free v7.1.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M224 248a120 120 0 1 0 0-240 120 120 0 1 0 0 240zm-29.7 56C95.8 304 16 383.8 16 482.3 16 498.7 29.3 512 45.7 512l356.6 0c16.4 0 29.7-13.3 29.7-29.7 0-98.5-79.8-178.3-178.3-178.3l-59.4 0z"/></svg></a></li>
        </ul>
    </div>
</nav>

<style lang="scss">
    nav {
        height: 40px;
        background-color: #fff;

        ul {
            display: flex;
            flex-direction: row;
            gap: 2rem;
            
            li {
                list-style: none;
                display: block;
                position: relative;
                height: 2rem;

                &:first-child {
                    margin-right: auto;
                }

                a {
                    line-height: 2rem;
                    svg {
                        padding-top: 0.4rem;
                        width: 1.2rem;
                    }   
                }
            }
        }
    }

</style>