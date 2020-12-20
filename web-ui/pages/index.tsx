import QuotesList from "components/QuotesList";
import { QuoteResponse } from "types";

interface IndexPageProps {
  data: QuoteResponse;
}

const IndexPage: React.FC<IndexPageProps> = ({ data }) => {
  return (
    <>
      <h2 className="text-center text-4xl text-indigo-900 font-display font-semibold lg:text-left xl:text-5xl xl:text-bold">
        Search for Quotes
        <div className="text-right text-sm text-indigo-400 font-italic">
          {data.totalHits} results in {data.milliseconds} ms
        </div>
        <div>
            <QuotesList quotes={data.quotes}></QuotesList>
        </div>
      </h2>
    </>
  );
};

export async function getServerSideProps() {
  const res = await fetch(`http://localhost:9000/search`);
  const data = await res.json();
  console.log(data);
  return { props: { data } };
}

export default IndexPage;
