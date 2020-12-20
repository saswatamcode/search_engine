import { Quote } from "types";

interface QuoteItemProps {
  quote: Quote;
}
const QuoteItem: React.FC<QuoteItemProps> = ({ quote }) => {
  return (
    <>
      <div className="mt-5 p-3 shadow rounded bg-blue-200">
        <div className="flex flex-col">
          <h3 className="text-left text-indigo-800 font-bold text-2xl">
            {quote.content}
          </h3>
          <h6 className="text-right text-indigo-800 font-bold text-2xl">
            by {quote.author}
          </h6>
        </div>
      </div>
    </>
  );
};

export default QuoteItem;
