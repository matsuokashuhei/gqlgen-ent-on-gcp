import { format, formatRFC3339, parse } from "date-fns";
import { useEffect, useState, VFC } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { Timetable } from "../../components/Timetable";
import {
  GetClassesBySchoolQuery,
  useGetClassesBySchoolLazyQuery,
} from "../../generated/graphql";

export const ClassesPage: VFC = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();

  const [getClassesBySchool, { data, loading, error }] =
    useGetClassesBySchoolLazyQuery();

  const [date, setDate] = useState<string>();
  const [school, setSchool] = useState<GetClassesBySchoolQuery["school"]>();
  const [studioId, setStudioId] = useState<string | null>();
  const [roomId, setRoomId] = useState<string | null>();

  useEffect(() => {
    const date = searchParams.get("date");
    // setStudioId(searchParams.get("studioId"));
    // setRoomId(searchParams.get("roomId"));
    if (!date) {
      const params = new URLSearchParams();
      params.append("date", date ?? format(new Date(), "yyyyMMdd"));
      // params.append("studioId", studioId ?? "");
      // params.append("roomId", roomId ?? "");
      navigate(`/classes?${params.toString()}`);
      return;
    }
    setDate(date);
  }, [searchParams, date, studioId, roomId, navigate]);

  useEffect(() => {
    if (!date) return;
    getClassesBySchool({
      variables: {
        id: "1",
        date: formatRFC3339(parse(date, "yyyyMMdd", new Date())),
      },
    });
  }, [date, getClassesBySchool]);

  useEffect(() => {
    if (!data) return;
    setSchool(data.school);
    setStudioId(searchParams.get("studioId") ?? data.school.studios[0].id);
    setRoomId(searchParams.get("roomId") ?? data.school.studios[0].rooms[0].id);
  }, [data, searchParams]);

  // useEffect(() => {
  //   if (school) {
  //   }
  // }, [school]);

  if (!date || !school || !studioId || !roomId) return <></>;

  return (
    <Timetable
      school={school}
      selectedStudioId={studioId}
      selectedRoomId={roomId}
      date={date}
      handleSelectStudioId={(studioId) => setStudioId(studioId)}
      handleSelectRoomId={(roomId) => setRoomId(roomId)}
      onClickClass={(clazz) => navigate(`/classes/${clazz.id}`)}
      onClickNewClass={(schedule) =>
        navigate(`/schedules/${schedule.id}/classes/new`)
      }
    />
  );
};
