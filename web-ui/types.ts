export interface Quote {
  content: string;
  author: string;
}

export interface QuoteResponse {
  milliseconds: number;
  totalHits: number;
  quotes: [Quote];
}
