# asuanna

A Go-based GUI utility for switching wallpapers and syncing desktop color schemes on Wayland / X11.

![Asuanna Interface](./image_2026-09-26_23-54-09.png)

Asuanna provides a graphical interface to quickly apply wallpapers, update color palettes via pywal, and restart desktop services (AGS, SwayNC) along with external app themes.

## Features

* **Image Preview:** View the selected wallpaper directly in the UI.
* **Wallpaper Switching:** Trigger any external daemon (`awww`,`feh`, etc.).
* **Color Schemes:** Integration with `pywal` and custom scripts (e.g., Telegram themes).
* **Service Management:** One-click restart for AGS and SwayNC.
* **Custom Profiles:** Load alternative JSON configuration files on the fly.

### Running from Source

```bash
# Clone the repository
git clone https://github.com/NyxAiko7/asuanna.git
cd asuanna

# Install dependencies and run
go run .

# Building the Binary
go build -o asuanna .
./asuanna
