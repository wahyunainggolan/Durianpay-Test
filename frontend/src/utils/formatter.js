export class Formatter {
  static currency(val) {
    if (val === null || val === undefined) return "-"

    return new Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR"
    }).format(val)
  }

  static date(date) {
    if (!date) return "-"

    return new Date(date).toLocaleString("id-ID", {
      day: "2-digit",
      month: "short",
      year: "numeric"
    })
  }
}