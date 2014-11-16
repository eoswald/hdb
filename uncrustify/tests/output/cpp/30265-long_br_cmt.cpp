
namespace a::b
{
   void foo::bar(int xx)
   {
      switch (xx)
      {
      case 1:
         // filler
         while (true)
         {
            if (something)
            {
               do_something();
            }
            else if (something_else)
            {
               do_something_else();
            }
            else
            {
               dont_do_anything();
               break;
            }
         }
         break;

      case 2:
         handle_two();

      default:
         handle_the_rest();
         break;
      } // switch
   } // foo::bar
} // namespace a::b
