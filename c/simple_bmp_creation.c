#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

// for Linux platform, plz make sure the size of data type is correct for BMP spec.
// if you use this on Windows or other platforms, plz pay attention to this.
typedef int LONG;
typedef unsigned char BYTE;
typedef unsigned int DWORD;
typedef unsigned short WORD;

// __attribute__((packed)) on non-Intel arch may cause some unexpected error, plz be informed.

typedef struct tagBITMAPFILEHEADER
{
	WORD	bfType; // 2  /* Magic identifier */
	DWORD   bfSize; // 4  /* File size in bytes */
	WORD	bfReserved1; // 2
	WORD	bfReserved2; // 2
	DWORD   bfOffBits; // 4 /* Offset to image data, bytes */ 
} __attribute__((packed)) BITMAPFILEHEADER;

typedef struct tagBITMAPINFOHEADER
{
	DWORD	biSize; // 4 /* Header size in bytes */
	LONG	 biWidth; // 4 /* Width of image */
	LONG	 biHeight; // 4 /* Height of image */
	WORD	 biPlanes; // 2 /* Number of colour planes */
	WORD	 biBitCount; // 2 /* Bits per pixel */
	DWORD	biCompress; // 4 /* Compression type */
	DWORD	biSizeImage; // 4 /* Image size in bytes */
	LONG	 biXPelsPerMeter; // 4
	LONG	 biYPelsPerMeter; // 4 /* Pixels per meter */
	DWORD	biClrUsed; // 4 /* Number of colours */ 
	DWORD	biClrImportant; // 4 /* Important colours */ 
} __attribute__((packed)) BITMAPINFOHEADER;

/*
typedef struct tagRGBQUAD
{
	unsigned char	rgbBlue;   
	unsigned char	rgbGreen;
	unsigned char	rgbRed;  
	unsigned char	rgbReserved;
} RGBQUAD;
* for biBitCount is 16/24/32, it may be useless
*/

typedef struct
{
		BYTE	b;
		BYTE	g;
		BYTE	r;
} RGB_data; // RGB TYPE, plz also make sure the order

int bmp_generator(char *filename, int width, int height, unsigned char *data)
{
	BITMAPFILEHEADER bmp_head;
	BITMAPINFOHEADER bmp_info;
	int size = width * height * 3;

	bmp_head.bfType = 0x4D42; // 'BM'
	bmp_head.bfSize= size + sizeof(BITMAPFILEHEADER) + sizeof(BITMAPINFOHEADER); // 24 + head + info no quad	
	bmp_head.bfReserved1 = bmp_head.bfReserved2 = 0;
	bmp_head.bfOffBits = bmp_head.bfSize - size;
	// finish the initial of head

	bmp_info.biSize = 40;
	bmp_info.biWidth = width;
	bmp_info.biHeight = height;
	bmp_info.biPlanes = 1;
	bmp_info.biBitCount = 24; // bit(s) per pixel, 24 is true color
	bmp_info.biCompress = 0;
	bmp_info.biSizeImage = size;
	bmp_info.biXPelsPerMeter = 0;
	bmp_info.biYPelsPerMeter = 0;
	bmp_info.biClrUsed = 0 ;
	bmp_info.biClrImportant = 0;
	// finish the initial of infohead;

	// copy the data
	FILE *fp;
	if (!(fp = fopen(filename,"wb"))) return 0;

	fwrite(&bmp_head, 1, sizeof(BITMAPFILEHEADER), fp);
	fwrite(&bmp_info, 1, sizeof(BITMAPINFOHEADER), fp);
	fwrite(data, 1, size, fp);
	fclose(fp);

	return 1;
}

BYTE to_byte(double d)
{
	if(d < 0)
		return 0;
	if(0xff < d)
		return 0xff;
	return (int)d;
}

int main(int argc, char **argv)
{
	const int img_w = 1280, img_h = 720;

	RGB_data buffer[img_h][img_w];
	memset(buffer, 0, sizeof(buffer));


	for (int i = 0; i < img_h; i++)
	{
		for (int j = 0; j < img_w; j++)
		{
			buffer[i][j].r = 0xff - to_byte(0xff);
			buffer[i][j].g = 0x00;
			buffer[i][j].b = 0x00;
		}
	}

	bmp_generator("./test.bmp", img_w, img_h, (BYTE*)buffer);	
	return EXIT_SUCCESS;
}


