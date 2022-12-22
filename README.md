<!-- PROJECT SHIELDS -->
<div align="center">

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

</div>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/HugoSohm/spotify-top">
    <img src="https://cdn.pixabay.com/photo/2021/12/11/06/40/spotify-6862049_1280.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Spotify Top</h3>

  <p align="center">
    Get your top artists and tracks from Spotify
    <br />
    <a href="https://github.com/HugoSohm/spotify-top"><strong>Explore the docs »</strong></a>
  </p>
</div>

<!-- GETTING STARTED -->
## Getting Started
### Available endpoints

- [**Login with spotify** - /login](https://api.spotifytop.hugosohm.fr/login)
- [**Get top artists** - /top/artists](https://api.spotifytop.hugosohm.fr/top/artists)
- [**Get top tracks** - /top/tracks](https://api.spotifytop.hugosohm.fr/top/tracks)

### Prerequisites
#### Install GoLang

```bash
brew install go
```

#### Add the following variables to you shell configuration
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

#### Run the following git config command
```bash
git config --global url.git@github.com:.insteadOf https://github.com/
```

### Installation
> Get free API key at [Spotify developer dashboard](https://developer.spotify.com/dashboard/login)

#### Clone the repository
```bash
git clone https://github.com/HugoSohm/spotify-top.git
```

#### Install packages
```sh
go mod tidy
```

#### Create env file and fill variables
```bash
cp .env.example .env
```

#### Run project
```sh
# Production
go run server.go

# Development
air
```

<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/HugoSohm/spotify-top.svg?style=for-the-badge
[contributors-url]: https://github.com/HugoSohm/spotify-top/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/HugoSohm/spotify-top.svg?style=for-the-badge
[forks-url]: https://github.com/HugoSohm/spotify-top/network/members
[stars-shield]: https://img.shields.io/github/stars/HugoSohm/spotify-top.svg?style=for-the-badge
[stars-url]: https://github.com/HugoSohm/spotify-top/stargazers
[issues-shield]: https://img.shields.io/github/issues/HugoSohm/spotify-top.svg?style=for-the-badge
[issues-url]: https://github.com/HugoSohm/spotify-top/issues
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/hugo-sohm