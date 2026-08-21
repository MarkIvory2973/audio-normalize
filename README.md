# Audio Normalize

Normalize the loudness of audio by FFmpeg.

## Installation

> **Note:**
>
> The following binaries are required and must be available in your `$PATH`:
>
> - [x] `ffmpeg`

### GitHub Releases

Download latest release from [GitHub Releases](https://github.com/MarkIvory2973/audio-normalize/releases/latest).

### Build from source

#### Requirements

- Go 1.26+
- nFPM
- UPX
- GNU Make
- Git

Clone the repository:

```bash
git clone https://github.com/MarkIvory2973/audio-normalize.git
cd audio-normalize
```

Install dependencies:

```bash
make install
```

Build binaries:

```bash
make build
```

Clean files:

```bash
make clean
```

## Usage

Run the following command:

```bash
./audio-normalize -help
```