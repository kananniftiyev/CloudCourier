import { useParams } from "react-router-dom";

function File() {
  const { uuid } = useParams();

  return (
    <>
      <h1>{uuid}</h1>
    </>
  );
}

export default File;
