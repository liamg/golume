# Golume

A tiny wrapper for addressing the active pulse audio output, hacked together for nicer i3 key mappings when switching between laptop speakers/headphones.

## Install

```
sudo make install
```

## Usage

```
  --change-volume int
       Change the volume by a given percent (negative to decrease volume) on all outputs
  --toggle-mute
       Toggle mute on all outputs
```

## Example i3 config

```
bindsym XF86AudioRaiseVolume exec --no-startup-id golume --change-volume 5  # increase sound volume
bindsym XF86AudioLowerVolume exec --no-startup-id golume --change-volume -5 # decrease sound volume
bindsym XF86AudioMute exec --no-startup-id golume --toggle-mute             # toggle mute sound
```
