typedef union {
	unsafe$Pointer ptr;
	complex128     c128;
	slice          sli;

	// Additional fields, useful for debugging:
	int8      i8;
	int16     i16;
	int32     i32;
	int64     i64;
	float32   f32;
	float64   f64;
	complex64 c64;
	string    str;
} ival;

typedef struct {
	ival val$;
	const void *itab$;
} interface;

#define EQUALI(lhs, rhs) ({                           \
	typeof(lhs) a = (lhs);                            \
	typeof(rhs) b = (rhs);                            \
	a.itab$ == b.itab$ && a.val$.c128 == b.val$.c128; \
})

#define CAST(t, e) ({                               \
	union {typeof(e) in; typeof(t) out;} cast = {}; \
	cast.in = (e);                                  \
	cast.out;                                       \
})

#define INTERFACE(e, itab) ((interface){CAST(ival, e), itab})
#define IVAL(i, typ) CAST(typ, (i).val$)
#define NILI ((interface){})
#define ISNILI(e) ((e).itab$ == nil)
