import { Quote } from "types";
import QuoteItem from "./QuoteItem";

interface QuoteListProps {
  quotes: [Quote];
}
const QuotesList: React.FC<QuoteListProps> = ({ quotes }) => {
  return (
    <>
      {quotes.map((quote, index) => (
        <QuoteItem key={index} quote={quote} />
      ))}
    </>
  );
};

export default QuotesList;
