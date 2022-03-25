import { PlusSmIcon } from "@heroicons/react/solid";
import { formatRFC3339, parse } from "date-fns";
import { useEffect, VFC } from "react";
import { Link } from "react-router-dom";
import { PartialDeep } from "type-fest";
import {
  GetClassesBySchoolQuery,
  useGetClassesBySchoolLazyQuery,
} from "../generated/graphql";

type Props = {
  date: string;
  onClickClass: (clazz: ClassType) => void;
  onClickNewClass?: (schedule: PartialDeep<SchedulesType[number]>) => void;
};

type StdiosType = GetClassesBySchoolQuery["school"]["studios"];
type StdioType = StdiosType[number];
type RoomsType = StdioType["rooms"];
type RoomType = RoomsType[number];
type SchedulesType = RoomType["schedules"];
type ClassType = NonNullable<SchedulesType[number]["class"]>;
type InstructorType = ClassType["instructor"];

export const Timetable: VFC<Props> = ({ date, onClickClass }) => {
  const [getClassesBySchool, { data, loading, error }] =
    useGetClassesBySchoolLazyQuery();

  useEffect(() => {
    getClassesBySchool({
      variables: {
        id: "1",
        date: formatRFC3339(parse(date, "yyyyMMdd", new Date())),
      },
    });
  }, [date, getClassesBySchool]);

  const renderStudios = (studios: StdiosType) => {
    return studios.map((studio) => renderStudio(studio));
  };

  const renderStudio = (studio: StdioType) => {
    return (
      <div key={studio.id}>
        {studio.name}
        {renderRooms(studio.rooms)}
      </div>
    );
  };

  const renderRooms = (rooms: RoomsType) => {
    return rooms.map((room) => renderRoom(room));
  };

  const renderRoom = (room: RoomType) => {
    return (
      <div key={room.id}>
        {room.name}
        {renderSchedules(room.schedules)}
      </div>
    );
  };

  const renderSchedules = (schedules: SchedulesType) => {
    return (
      <div className="grid grid-cols-8">
        {[0, 1, 2, 3, 4, 5, 6, 7].map((dayOfWeek) => {
          return <div key={dayOfWeek}>{dayOfWeek}</div>;
        })}
        {renderTimeslots(schedules)}
      </div>
    );
  };

  const renderTimeslots = (schedules: SchedulesType) => {
    return [
      "13:00",
      "14:15",
      "15:30",
      "16:45",
      "18:00",
      "19:15",
      "20:30",
      "21:45",
    ].map((startTime) => {
      return [0, 1, 2, 3, 4, 5, 6, 7].map((dayOfWeek) => {
        if (dayOfWeek === 0) {
          return <div key={dayOfWeek}>{startTime}</div>;
        }
        const schedule = schedules.find(
          (schedule) =>
            schedule.dayOfWeek === dayOfWeek && schedule.startTime === startTime
        );
        if (!schedule) return <div key={`${dayOfWeek}-none`}></div>;
        if (schedule.class) {
          return (
            <div key={`${dayOfWeek}-${schedule.id}`}>
              {renderClass(schedule.class)}
            </div>
          );
        } else {
          return (
            <div key={`${dayOfWeek}-${schedule.id}`}>
              {renderNewClassButton(schedule)}
            </div>
          );
        }
      });
    });
  };

  const renderClass = (clazz: ClassType) => {
    if (!clazz) return <></>;
    return (
      <div onClick={() => onClickClass(clazz)}>
        <div>{clazz.name}</div>
        <div>{clazz.level}</div>
        {renderInstructor(clazz.instructor)}
      </div>
    );
  };

  const renderInstructor = (instructor: InstructorType) => {
    return <div>{instructor.name}</div>;
  };

  const renderNewClassButton = (
    schedule?: PartialDeep<SchedulesType[number]>
  ) => {
    if (!schedule) return <></>;
    return (
      <Link to={`/schedules/${schedule.id}/classes/new`}>
        <PlusSmIcon className="h-4 w-4" />
      </Link>
    );
  };

  if (!data) {
    return <></>;
  }
  const {
    school: { studios },
  } = data;

  return (
    <>
      <h1>クラス</h1>
      {renderStudios(studios)}
    </>
  );
};
