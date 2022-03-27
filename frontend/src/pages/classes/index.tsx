import { format } from "date-fns";
import { useEffect, useState, VFC } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { Timetable } from "../../components/Timetable";

export const ClassesPage: VFC = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();

  const [date, setDate] = useState<string | null>(null);

  useEffect(() => {
    const date = searchParams.get("date");
    if (!date) {
      navigate(`/classes?date=${format(new Date(), "yyyyMMdd")}`);
      return;
    }
    setDate(date);
  }, [searchParams, navigate]);

  if (!date) return <></>;

  return (
    <Timetable
      date={date}
      onClickClass={(clazz) => navigate(`/classes/${clazz.id}`)}
      onClickNewClass={(schedule) =>
        navigate(`/schedules/${schedule.id}/classes/new`)
      }
    />
  );
};
