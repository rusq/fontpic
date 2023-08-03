# fontpic

Goal of the project - allow to easily generate images with text using widely
available CGA/EGA/VGA ROM fonts.  This can be useful for generating captcha
images for website or telegram chats.

It comes with 8x8, 8x14 and 8x16 fonts in DOS 866 (Cyrillic) encoding in
[fnt](/fnt) directory.  Font files 08x08.fnt, 08x14.fnt, 08x16.fnt were
distributed with the freeware [KeyRus program][1], and are used in this
library by default as a tribute to Dmitry Gurtyak (1971-1998), author of
KeyRus.

Currenly the project supports only raw font files, it's easy to tell if it's
a raw FNT by looking at the file size:

- 8x8 - 2048 bytes
- 8x14 - 3584 bytes
- 8x16 - 4096 bytes

## Where to get more fonts

1. There is a great project that contains a lot of fonts extracted from
   different ROMs: [romfont][2]
2. Extract fonts from a BIOS of the old PC.  Read the [romfont][2] repository
   README.
3. Unpack fonts from the Abandonware programs.  I.e. DOS distribution includes
   '*.CPI' files that contain fonts.  You can use [psf2inc][3] utility to
   extract them.  If you went down that path, you'd probably know what to do
   with the output.
4. Convert BDF fonts.


## Licensing

BSD 3-clause.  See [LICENSE](/LICENSE).

Included fonts are freeware (c) Dmitry Gurtyak.


[1]: https://en.wikipedia.org/wiki/KeyRus
[2]: https://github.com/spacerace/romfont
[3]: https://www.mankier.com/1/psf2inc
