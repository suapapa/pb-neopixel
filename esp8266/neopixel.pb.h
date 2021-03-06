/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.2-dev */

#ifndef PB_NEOPIXEL_PB_H_INCLUDED
#define PB_NEOPIXEL_PB_H_INCLUDED
#include "pb.h"

#if PB_PROTO_HEADER_VERSION != 40
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Struct definitions */
typedef struct _LED {
    uint32_t index;
    uint32_t color;
} LED;

typedef struct _NeoPixel {
    bool clear;
    pb_callback_t strip;
} NeoPixel;


/* Initializer values for message structs */
#define LED_init_default                         {0, 0}
#define NeoPixel_init_default                    {0, {{NULL}, NULL}}
#define LED_init_zero                            {0, 0}
#define NeoPixel_init_zero                       {0, {{NULL}, NULL}}

/* Field tags (for use in manual encoding/decoding) */
#define LED_index_tag                            1
#define LED_color_tag                            2
#define NeoPixel_clear_tag                       1
#define NeoPixel_strip_tag                       2

/* Struct field encoding specification for nanopb */
#define LED_FIELDLIST(X, a) \
X(a, STATIC,   REQUIRED, UINT32,   index,             1) \
X(a, STATIC,   REQUIRED, UINT32,   color,             2)
#define LED_CALLBACK NULL
#define LED_DEFAULT NULL

#define NeoPixel_FIELDLIST(X, a) \
X(a, STATIC,   REQUIRED, BOOL,     clear,             1) \
X(a, CALLBACK, REPEATED, MESSAGE,  strip,             2)
#define NeoPixel_CALLBACK pb_default_field_callback
#define NeoPixel_DEFAULT NULL
#define NeoPixel_strip_MSGTYPE LED

extern const pb_msgdesc_t LED_msg;
extern const pb_msgdesc_t NeoPixel_msg;

/* Defines for backwards compatibility with code written before nanopb-0.4.0 */
#define LED_fields &LED_msg
#define NeoPixel_fields &NeoPixel_msg

/* Maximum encoded size of messages (where known) */
#define LED_size                                 12
/* NeoPixel_size depends on runtime parameters */

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif
