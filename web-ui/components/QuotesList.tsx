import { Quote } from "types";
import QuoteItem from "./QuoteItem";

interface QuoteListProps {
  quotes: [Quote];
}
const QuotesList: React.FC<QuoteListProps> = ({ quotes }) => {
  return (
    <>
      {quotes.map((quote) => (
        <QuoteItem quote={quote} />
      ))}
    </>
  );
};

export default QuotesList;